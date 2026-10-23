
repositories = {
    ["mq"] = {
        url = "https://codeberg.org/menos/mq"
    },
    ["mbm"] = {
        url = "https://codeberg.org/menos/mbm"
    },
    ["nvim"] = {
        url = "https://github.com/neovim/neovim",
        build = {
            system = "make",
            args = {"-CMAKE_BUILD=Release"}
        }
    },
    ["suckless"] = {
        base = true,
        url = "https://git.suckless.org/",
    }
}

opts = {
    systems = {
        ["make"] = {
            build_cmd = function()
                os.execute("make");
            end,
            install_cmd = function()
                os.execute("make install");
            end,
            root_markers = {"Makefile"}
        }
    }
}

packages = {
    "mbm",
    "mq",
    "nvim",
    ["dmenu"] = {
        base = "suckless",
        path = "dmenu"
    }
}

