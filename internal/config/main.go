package config

import (
	"errors"
	"iko/internal/build"

	lua "github.com/yuin/gopher-lua"
)

type RepositoryBuild struct {
	System string
	Executable string
	Args []string
}

type Package struct {
	Name string
	Base string
	Path string
}

type Repository struct {
	Name string
}

type Pkg struct {
	Name string
	Repository string
	Build build.BuildSystem
	Base string
	Path string
}

func GetState() (*lua.LState, error) {
	l := lua.NewState();
	err := l.DoFile("tests/init.lua");
	if err != nil {
		return l, err;
	}
	return l, nil;
}

func ReadRepositories() ([]Repository, error) {
	l, err := GetState();
	if err != nil {
		return nil, err;
	}
	defer l.Close();

	var repos []Repository;

	table, ok := l.GetGlobal("repositories").(*lua.LTable);
	if !ok {
		return nil, errors.New(`repositories should be a table`);
	}

	table.ForEach(func (k, v lua.LValue) {
		var repo Repository;


		repos = append(repos, repo);
	})

	return nil, nil;
}

func ReadPackages() ([]Package, error) {
	l, err := GetState();
	if err != nil {
		return nil, err;
	}
	defer l.Close();

	var pkgs []Package;
	pkgValue := l.GetGlobal("packages");
	pkgTable, ok := pkgValue.(*lua.LTable);
	if !ok {
		return nil, errors.New(`packages should be a table`);
	}

	pkgTable.ForEach(func(k, v lua.LValue) {
		var pkg Package;

		switch k.Type() {
		case lua.LTNumber:
			pkg.Name = v.String()
		case lua.LTString:
			pkg.Name = k.String();
		}

		vt, ok := v.(*lua.LTable);
		if !ok {
			pkgs = append(pkgs, pkg);
			return;
		}
		vt.ForEach(func(k, v lua.LValue) {
			pkg.Base = l.GetField(vt, "base").String();
			pkg.Path = l.GetField(vt, "path").String();
		})
		pkgs = append(pkgs, pkg);
	})
	return pkgs, nil;
}

