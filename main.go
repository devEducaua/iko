package main

import (
	"errors"
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
		helpF bool;
		installF bool;
		removeF bool;
		syncF bool;
		updateF bool;
		listF bool;
		onlyF string;
		excludeF string;
	)
	f.Var("install", "i", &installF);
	f.Var("remove", "r", &removeF);
	f.Var("update", "u", &updateF);
	f.Var("sync", "s", &syncF);
	f.Var("list", "l", &listF);
	f.Var("only", "o", &onlyF);
	f.Var("exclude", "x", &excludeF);
	f.Var("help", "h", &helpF);

	if err := f.Parse(argv); err != nil {
		return err;
	}
	if len(argv) < 1 {
		// f.ShowHelp();
		return nil;
	}

	switch {
	case installF:
		err := commands.Install(onlyF, excludeF);
		if err != nil {
			return err;
		}
	case removeF:
		err := commands.Remove(onlyF, excludeF);
		if err != nil {
			return err;
		}
	case updateF:
		err := commands.Update(onlyF, excludeF);
		if err != nil {
			return err;
		}
	case syncF:
		err := commands.Sync(onlyF, excludeF);
		if err != nil {
			return err;
		}
	case listF:
		err := commands.List();
		if err != nil {
			return err;
		}
	case helpF:
		// f.ShowHelp()
	default:
		// f.ShowHelp()
	}
	return nil;
}

func DetectBuildSystem(buildMarkers map[string]string, path string) (string, error) {
	entries, err := os.ReadDir(path);
	if err != nil {
		return "", err;
	}

	for _,e := range entries {
		if sys, ok := buildMarkers[e.Name()]; ok {
			return sys, nil;
		}
	}
	return "", errors.New("could not find build system");
}
