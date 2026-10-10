package main

import (
	"fmt"
	"iko/internal/ui"
	"os"
)

type Pkg struct {
	Name string
	Repo string
	Url string
}

type Repo struct {
	Name string
	Url string
	Base bool
}

func main() {
	// repos := []Repo{
	// 	Repo{
	// 		Name: "nvim",
	// 		Url: "https://github.com/neovim/neovim",
	// 	},
	// 	Repo{
	// 		Name: "github",
	// 		Url: "https://github.com",
	// 		Base: true,
	// 	},
	// }

	dat, err := os.ReadFile("syntax.ui");
	if err != nil {
		panic(err);
	}

	p, err := ui.Parse(string(dat));
	if err != nil {
		panic(err);
	}
	fmt.Println(p.Entries);
}

