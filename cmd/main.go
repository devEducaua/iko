package main

import "ui/internal/lang"

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

	err :=  lang.Parse("syntax.ui");
	if err != nil {
		panic(err);
	}
}

