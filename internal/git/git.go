package git

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// returned bool means alreadyClonned
func ClonePkg(name, url, configPath string) (bool, error) {
	path := filepath.Join(configPath, name);
	out, err := gitCommand("", "clone", url, path);
	if err != nil {
		if strings.HasPrefix(err.Error(), "exit status 128") {
			return true, nil;
		}
		return false, err;
	}
	fmt.Printf("[GIT]: cloning %v into %v from %v: %v\n", name, path, url, out);
	return false, nil;
}

func gitCommand(dir string, command ...string) (string, error) {
	cmd := exec.Command("git", command...);
	cmd.Dir = dir;
	bt, err := cmd.CombinedOutput();
	if err != nil {
		return "", err;
	}
	return string(bt), nil;
}

func Pull(name, path string) (bool, error) {
	// TODO: branch and commits options later
	out, err := gitCommand(path, "pull");
	if err != nil {
		return false, err;
	}
	if strings.HasPrefix(out, "Already up to date.") {
		return false, nil;
	}
	fmt.Printf("[PULL]: %v\n", name);
	return true, nil;
}

