package git

import (
	"os"
	"os/exec"
	"path/filepath"
)

func ClonePkg(url, name string) error {
	path, err := getOutPath(name);
	if err != nil {
		return err;
	}

	if _, err := gitCommand("clone", url, path); err != nil {
		return err;
	}

	return nil;
}

func gitCommand(command ...string) (string, error) {
	cmd := exec.Command("git", command...);

	bt, err := cmd.Output();
	if err != nil {
		return "", err;
	}
	return string(bt), nil;
}

func gitPull(name string) error {
	// TODO: go to directory
	out, err := gitCommand("pull");
	if err != nil {
		return err;
	}



	return nil;
}

func getOutPath(name string) (string, error) {
	xdg := os.Getenv("XDG_CONFIG_HOME");
	if xdg == "" {
		home, err := os.UserConfigDir();
		if err != nil {
			return "", err;
		}
		xdg = filepath.Join(home, ".config");
	}

	ikoDir := filepath.Join(xdg, "iko");
	if err := os.MkdirAll(uyDir, 0755); err != nil {
		return "", err;
	}
	return filepath.Join(uyDir, name), nil;
}

