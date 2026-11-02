package main

import (
	"fmt"
	"iko/internal/commands"
	"iko/internal/flags"
	"os"
)

func main() {
	argv := os.Args[1:];
	if err := parseArgs(argv); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err);
		os.Exit(1);
	}
}

func parseArgs(argv []string) error {
	f := make(flags.Flags);
	var (
		onlyF string;
		excludeF string;
	);
	
	installF := f.VarBool("install", "i", "install listed pkgs");
	removeF := f.VarBool("remove", "r", "remove not declared pkgs");
	updateF := f.VarBool("update", "u", "update pkgs");
	syncF := f.VarBool("sync", "s", "remove, install and update pkgs");
	listF := f.VarBool("list", "l", "list installed pkgs");
	helpF := f.VarBool("help", "h", "show this page");
	f.Var("only", "o", &onlyF, "only apply action on those pkgs");
	f.Var("exclude", "x", &excludeF, "exclude pkgs from action");

	if err := f.Parse(argv); err != nil {
		return err;
	}
	if len(argv) < 1 {
		f.Usage();
		return nil;
	}

	switch {
	case *installF:
		err := commands.Install(onlyF, excludeF);
		if err != nil {
			return err;
		}
	case *removeF:
		err := commands.Remove(onlyF, excludeF);
		if err != nil {
			return err;
		}
	case *updateF:
		err := commands.Update(onlyF, excludeF);
		if err != nil {
			return err;
		}
	case *syncF:
		err := commands.Sync(onlyF, excludeF);
		if err != nil {
			return err;
		}
	case *listF:
		err := commands.List();
		if err != nil {
			return err;
		}
	case *helpF:
		f.Usage();
	default:
		f.Usage();
	}
	return nil;
}

