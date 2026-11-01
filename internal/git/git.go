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

func gitPull(name string) (bool, error) {
	// TODO: to implement
	return false, nil;
}

func getConfigPath() (string, error) {
	xdg := os.Getenv("XDG_CONFIG_HOME");
	if xdg == "" {
		home, err := os.UserConfigDir();
		if err != nil {
			return "", err;
		}
		xdg = filepath.Join(home, ".config");
	}

	ikoDir := filepath.Join(xdg, "iko");
	if err := os.MkdirAll(ikoDir, 0755); err != nil {
		return "", err;
	}
	return ikoDir, nil;
}

