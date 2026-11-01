package pkgs

import (
	"fmt"
	"iko/internal/config"
	"iko/internal/git"
	"os"
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)


type Pkg struct {
	Name string
	System config.System
	Url string
	Repository string
	Path string
}

func GetPkgs() (map[string]Pkg, *lua.LState, error) {
	l, err := config.GetState();
	if err != nil {
		return nil, nil, err;
	}
	packages, err := config.ReadPackages(l);
	if err != nil {
		return nil, nil, err;
	}
	repositories, err := config.ReadRepositories(l);
	if err != nil {
		return nil, nil, err;
	}
	systems, err := config.ReadSystems(l);
	if err != nil {
		return nil, nil, err;
	}

	pkgs := make(map[string]Pkg);
	for i := range packages {
		p := packages[i];

		var pkg Pkg;
		pkg.Name = p.Name;

		pkg.Repository = p.Name;
		if p.Base != "" {
			pkg.Repository = p.Base;
		}

		repo := repositories[pkg.Repository];
		pkg.System = systems[repo.System.Name];

		pkg.Url = repo.Url;
		if p.Base != "" {
			pkg.Url = fmt.Sprintf("%v%v", repo.Url, p.Path);
		}

		dir, err := config.GetPackagesDir();
		if err != nil {
			return nil, nil, err;
		}
		pkg.Path = filepath.Join(dir, pkg.Name);

		pkgs[pkg.Name] = pkg;
	}
	return pkgs, l, nil;
}

func FilterPkgs(pkgs map[string]Pkg, onlyArgs, excludeArgs []string) (map[string]Pkg, error) {
	if onlyArgs != nil {
		new := make(map[string]Pkg);
		for _,n := range onlyArgs {
			p, ok := pkgs[n];
			if !ok {
				return nil, fmt.Errorf("`%v` is not a known package name", n);
			}
			new[p.Name] = p;
		}
		return new, nil;
	}
	if excludeArgs != nil {
		for _,n := range excludeArgs { // 2
			_, ok := pkgs[n];
			if ok {
				delete(pkgs, n);
			}
		}
	}
	return pkgs, nil;;
}

func IsListed(pkgName string) (bool, error) {
	path, err := config.GetPackagesDir();
	if err != nil {
		return false, err;
	}
	entries, err := os.ReadDir(path);
	if err != nil {
		return false, err;
	}
	for _,e := range entries {
		if e.Name() == pkgName {
			return false, nil;
		}
	}
	return true, nil;
}

func Remove(pkg Pkg) error {
	dir, err := config.GetPackagesDir();
	if err != nil {
		return err;
	}
	if err := os.Chdir(filepath.Join(dir, pkg.Name)); err != nil {
		return err;
	}
	fmt.Printf("[REMOVE]: %v\n", pkg);
	err = pkg.System.Uninstall();
	return err;
}

func BuildInstall(pkg Pkg) error {
	dir, err := config.GetPackagesDir();
	if err != nil {
		return err;
	}
	alreadyClonned, err := git.ClonePkg(pkg.Name, pkg.Url, dir);
	if err != nil {
		return err;
	}
	if alreadyClonned { return nil; }

	if err := os.Chdir(filepath.Join(dir, pkg.Name)); err != nil {
		return err;
	}
	fmt.Printf("[BUILD]: %v from %v\n", pkg.Name, pkg.Url);
	err = pkg.System.Build();
	if err != nil {
		return err;
	}
	fmt.Printf("[INSTALL]: %v\n", pkg.Name);
	err = pkg.System.Install();
	return err;
}

