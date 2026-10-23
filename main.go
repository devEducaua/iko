package main

import (
	"fmt"
	"iko/internal/config"
	"os"
)

func main() {
	pkgs, err := config.ReadPackages();
	if err != nil {
		panic(err);
	}
	for i,p := range pkgs {
		fmt.Printf("[%v] %v\n", i, p);
	}
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
	return "", nil;
}
