
local home = os.getenv("HOME");
local prefix = home.."/.local"

repositories = {
    ["mew"] = {
        url = "https://codeberg.org/sewn/mew",
        system = {
            name = "make"
        }
    },
    ["nvim"] = {
        url = "https://github.com/neovim/neovim",

        system = {
            name = "make",
            build_args = {"-CMAKE_BUILD=Release"},
        },
    },
    ["suckless"] = {
        base = true,
        url = "https://git.suckless.org/",
        -- TODO: later it should automatically detect the build system
        system = {
            name = "make"
        }
    }
}

systems = {
    ["make"] = {
        build_cmd = function()
            os.execute("make");
        end,
        install_cmd = function()
            os.execute("make install PREFIX="..prefix);
        end,
        uninstall_cmd = function()
            os.execute("make uninstall PREFIX="..prefix);
        end,
        root_markers = {"Makefile", "makefile", "config.mk"}
    }
}

packages = {
    "mew",
    -- "nvim",
   ["dmenu"] = {
        base = "suckless",
        path = "dmenu"
    }
}

