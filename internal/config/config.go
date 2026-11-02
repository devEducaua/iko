package config

import (
	"errors"
	"fmt"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

type RepositorySystem struct {
	Name string
	// TODO; maybe change to []string
	BuildArgs string
	InstallArgs string
	UninstallArgs string
}

type Package struct {
	Name string
	Base string
	Path string
}

type Repository struct {
	Name string
	Url string
	System RepositorySystem
}

type BuildFunc func() error

type System struct {
	Name string
	Build BuildFunc
	Install BuildFunc
	Uninstall BuildFunc
	RootMarkers []string
}

func GetState() (*lua.LState, error) {
	f, err := GetConfigFile();
	if err != nil {
		return nil, err;
	}
	l := lua.NewState();
	err = l.DoFile(f);
	if err != nil {
		return l, err;
	}
	return l, nil;
}

func CollectBuildSystemMarkers(l *lua.LState, systems map[string]System) (map[string]string, error) {
	m := make(map[string]string);

	for _,s := range systems {
		for _,marker := range s.RootMarkers {
			m[marker] = s.Name;
		}
	}
	return m, nil;
}

func ReadRepositories(l *lua.LState) (map[string]Repository, error) {
	repos := make(map[string]Repository);

	table, ok := l.GetGlobal("repositories").(*lua.LTable);
	if !ok {
		return nil, errors.New(`repositories should be a table`);
	}

	table.ForEach(func (k, v lua.LValue) {
		var repo Repository;

		if k.Type() != lua.LTString {
			// TODO: handle error
			return;
		}
		repo.Name = k.String();

		vt, ok := v.(*lua.LTable);
		if !ok {
			// TODO: handle error
			return;
		}
		repo.Url = l.GetField(vt, "url").String();

		st, ok := l.GetField(vt, "system").(*lua.LTable);
		if ok {
			repo.System.Name = l.GetField(st, "name").String();
		}
		repos[repo.Name] = repo;
	})

	return repos, nil;
}

func ReadSystems(l *lua.LState) (map[string]System, error) {
	systems := make(map[string]System);
	sysValue := l.GetGlobal("systems");
	sysTable, ok := sysValue.(*lua.LTable);
	if !ok {
		return nil, errors.New("`systems` should be a table");
	}

	sysTable.ForEach(func(k, v lua.LValue) {
		var sys System;

		if k.Type() != lua.LTString {
			// TODO: handle error
			return;
		}
		sys.Name = k.String();

		vt, ok := v.(*lua.LTable);
		if !ok {
			// TODO: handle error
			return;
		}
		bcmd, err := GetFunctionFromTable(l, vt, "build_cmd");
		if err != nil {
			fmt.Printf("[ERROR]: %v", err)
			return
		}

		icmd, err := GetFunctionFromTable(l, vt, "install_cmd");
		if err != nil {
			fmt.Printf("[ERROR]: %v", err)
			return
		}

		ucmd, err := GetFunctionFromTable(l, vt, "uninstall_cmd");
		if err != nil {
			fmt.Printf("[ERROR]: %v", err)
			return
		}

		sys.Build = LuaToGo(l, bcmd);
		sys.Install = LuaToGo(l, icmd);
		sys.Uninstall = LuaToGo(l, ucmd);

		l.GetField(vt, "root_markers").(*lua.LTable).ForEach(func (k,v lua.LValue) {
			sys.RootMarkers = append(sys.RootMarkers, v.String()); 
		})

		systems[sys.Name] = sys;
	})
	return systems, nil;
}

func GetFunctionFromTable(l *lua.LState, table *lua.LTable, name string) (*lua.LFunction, error) {
	fn, ok := l.GetField(table, name).(*lua.LFunction);
	if !ok {
		return nil, fmt.Errorf("`%v` expected to be a function", name);
	}

	if fn == lua.LNil {
		return nil, fmt.Errorf("`%v` is a nil function", name);
	}

	return fn, nil;
}

func LuaToGo(l *lua.LState, fn *lua.LFunction) func () error {
	return func () error {
		return l.CallByParam(lua.P{
			Fn: fn,
			NRet: 0,
			Protect: true,
		});
	}
}

func ReadPackages(l *lua.LState) ([]Package, error) {
	var pkgs []Package;
	pkgValue := l.GetGlobal("packages");
	pkgTable, ok := pkgValue.(*lua.LTable);
	if !ok {
		return nil, errors.New("`packages` should be a table");
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
			if pkg.Base != "" && pkg.Path == "nil" {
				pkg.Path = pkg.Name;	
			}
		})
		pkgs = append(pkgs, pkg);
	})
	return pkgs, nil;
}

func GetOnlyExcludeArgs(onlyS, excludeS string) ([]string, []string) {
	var only, exclude []string = nil, nil;
	if splitted := strings.Split(onlyS, ","); len(splitted) > 0 && splitted[0] != ""{
		only = splitted;
	}
	if splitted := strings.Split(excludeS, ","); len(splitted) > 0 && splitted[0] != "" {
		exclude = splitted;
	}
	return only, exclude;
}

