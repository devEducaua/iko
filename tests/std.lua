
opts = {
    systems = {
        ["make"] = {
            build_cmd = function()
                os.execute("make");
            end,
            install_cmd = function()
                os.execute("make install");
            end,
            root_markers = {"Makefile", "GNUMakefile"}
        },
        ["cargo"] = {
            root_markers = {"Cargo.toml", "Cargo.lock"},
            build_cmd = function()
                os.execute("cargo build --release")
            end,
            install_cmd = function()
                os.execute("cargo install --path .") -- TODO: handle this dot.
            end
        },
        ["meson"] = {
            root_markers = {"meson.build"},
            build_cmd = function()
                os.execute("meson setup build")
                os.execute("ninja -C build")
            end,
            install_cmd = function()
                os.execute("ninja -C build install")
            end
        },
    }
}
