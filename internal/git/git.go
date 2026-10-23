package git

import (
	"os"
	"os/exec"
	"path/filepath"
)

func ClonePkg(url, name string) (string, error) {
	configPath, err := getConfigPath();
	if err != nil {
		return "", err;
	}
	path := filepath.Join(configPath, name);
	if _, err := gitCommand("clone", url, path); err != nil {
		return "", err;
	}
	return path, nil;
}

func gitCommand(command ...string) (string, error) {
	cmd := exec.Command("git", command...);

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

