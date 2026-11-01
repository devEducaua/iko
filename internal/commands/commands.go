package commands

import (
	"fmt"
	"iko/internal/config"
	"iko/internal/git"
	"iko/internal/pkgs"
	"os"
	"path/filepath"
)

func Sync(onlyArgs, excludeArgs string) error {
	err := Remove(onlyArgs, excludeArgs);
	if err != nil {
		return err;
	}
	err = Install(onlyArgs, excludeArgs);
	if err != nil {
		return err;
	}
	return Update(onlyArgs, excludeArgs);
}

func Remove(onlyArgs, excludeArgs string) error {
	// TODO: revise only and exclude in this command.
	
	dir, err := config.GetPackagesDir();
	if err != nil {
		return err;
	}

	entries, err := os.ReadDir(dir);
	if err != nil {
		return err;
	}

	pkgMap, l, err := pkgs.GetFilteredPkgs(onlyArgs, excludeArgs);
	if err != nil {
		return err;
	}
	defer l.Close();

	// TODO: should have the ability to remove not listed pkgs.

	for _,ent := range entries {
		if p, ok := pkgMap[ent.Name()]; ok {
			if err := os.Chdir(filepath.Join(dir, p.Name)); err != nil {
				return err;
			}
			fmt.Printf("[REMOVE]: `%v`\n", p.Name);

			sys, err := pkgs.GetBuildSystem(p);
			if err != nil {
				return err;
			}
			p.System = sys;

			err = p.System.Uninstall();
			if err != nil {
				return err;
			}
			continue;
		}
		fmt.Printf("[REMOVE]: failed to remove `%v`, no known way of uninstalling\n", ent.Name());
	}
	return nil;
}

func Update(onlyArgs, excludeArgs string) error {
	m, l, err := pkgs.GetFilteredPkgs(onlyArgs, excludeArgs);
	if err != nil {
		return err;
	}
	defer l.Close();

	for _,p := range m {
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

func Install(onlyArgs, excludeArgs string) error {
	m, l, err := pkgs.GetFilteredPkgs(onlyArgs, excludeArgs);
	if err != nil {
		return err;
	}
	defer l.Close();

	for _,p := range m {
		err = pkgs.BuildInstall(p);
		if err != nil {
			return err;
		}
	}
	return nil;
}

func List() error {
	m, l, err := pkgs.GetPkgs();
	if err != nil {
		return err;
	}
	defer l.Close();

	for _,p := range m {
		fmt.Printf("%v\n", p.Name);
	}

	return nil;
}


