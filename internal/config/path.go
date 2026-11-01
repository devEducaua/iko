package config

import (
	"os"
	"path/filepath"
)

func GetConfigPath() (string, error) {
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

func GetConfigFile() (string, error) {
	cp, err := GetConfigPath();
	if err != nil {
		return "", err;
	}
	return filepath.Join(cp, "init.lua"), nil;
}

func GetPackagesDir() (string, error) {
	cp, err := GetConfigPath();
	if err != nil {
		return "", err;
	}
	return filepath.Join(cp, "packages/"), nil;
}

