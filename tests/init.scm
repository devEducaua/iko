
(repositories 
    ("mq" (url "https://codeberg.org/menos/mq"))
    ("nvim"
        (url "https://github.com/neovim/neovim")
        (build ((system "make") (args ("-CMAKE_BUILD=Release")))))
    ("suckless" (base #t) (url "https://git.suckless.org/")))

(configure-system 
    "make"
    (build_cmd "make")
    (install_cmd "make")
    (root_markers ("Makefile")))

(packages
    '(mbm
      mq
      nvim))

