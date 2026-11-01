package commands

import (
	"fmt"
	"iko/internal/config"
	"iko/internal/git"
	"iko/internal/pkgs"
	"os"
)

func Sync(includeArgs, excludeArgs string) error {
	err := Remove(includeArgs, excludeArgs);
	if err != nil {
		return err;
	}
	err = Install(includeArgs, excludeArgs);
	if err != nil {
		return err;
	}
	return Update(includeArgs, excludeArgs);
}

func Remove(includeArgs, excludeArgs string) error {
	pkgMap, l, err := pkgs.GetPkgs();
	if err != nil {
		return err;
	}
	defer l.Close();

	dir, err := config.GetPackagesDir();
	if err != nil {
		return err;
	}

	entries, err := os.ReadDir(dir);
	if err != nil {
		return err;
	}

	for _,e := range entries {
		if p, ok := pkgMap[e.Name()]; !ok {
			err = pkgs.Remove(p);
			if err != nil {
				return err;
			}
		}
	}
	return nil;
}

func Update(includeArgs, excludeArgs string) error {
	pkgMap, l, err := pkgs.GetPkgs();
	if err != nil {
		return err;
	}
	defer l.Close();

	for _,p := range pkgMap {
		updated, err := git.Pull(p.Name, p.Path);
		if err != nil {
			// TOOD: should not break in the middle of the loop, just log
			return err;
		}
		if !updated {
			fmt.Printf("[UPDATE]: `%v` is already updated\n", p.Name);
			continue;
		}

		err = pkgs.BuildInstall(p);
		if err != nil {
			return err;
		}
	}
	return nil;
}

func Install(includeArgs, excludeArgs string) error {
	list, l, err := pkgs.GetPkgs();
	if err != nil {
		return err;
	}

	inc, ex := config.GetOnlyExcludeArgs(includeArgs, excludeArgs);
	filt, err := pkgs.FilterPkgs(list, inc, ex);
	if err != nil {
		return err;
	}

	defer l.Close();
	for _,p := range filt {
		err = pkgs.BuildInstall(p);
		if err != nil {
			return err;
		}
	}
	return nil;
}

