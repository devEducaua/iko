package pkgs

import (
	"iko/internal/config"
)

func GetPkgList() []config.Pkg {
	config.ReadPackages();
}

