# gobrew

Count all programs written/built in Go/Rust/Python/.. and distributed via Homebrew.

## why I built this?

I was curious. Are people _actually_ using Go more than Rust? Are there more software written with cmake or meson or ninja .. ?

Too many curious questions I wanted to answer them using _realworld statistics_.

I wrote a post about my findings regarding [Go vs Rust usage in Homebrew Core formulae on my website](https://abanoubhanna.com/posts/go-vs-rust-use-production/).

## commands

### build CLI app from source

```sh
# get deps, build the "gobrew" executable
$ go mod tidy && go build -o gobrew .
```

### count all packages that use a specific language

```sh
# will show the count of packages written in Go in Homebrew Core (by default)
$ ./gobrew
No language nor build system nor library is specified. Counting packages built in Go (by default):
974

$ ./gobrew -l go
974

$ ./gobrew --lang go
974

$ ./gobrew -l rust
524

$ ./gobrew -l cython
12

$ ./gobrew -l ruby
27

$ ./gobrew -l lua
51

$ ./gobrew -l zig
6
```

#### all versions of the specified language

- all versions of python

```sh
$ ./gobrew -l python
771
```

- specific version of python

```sh
$ ./gobrew -l python@3.12
718

$ ./gobrew -l python@3.11
67

$ ./gobrew -l python@3.10
10

$ ./gobrew -l python@3.9
6

$ ./gobrew -l python@3.8
2
```

### count all packages that use a specific build system or library

```sh
# get the packages built with "ninja" build system
$ ./gobrew -l ninja
253

$ ./gobrew -l gcc
75

$ ./gobrew -l cmake
1011

$ ./gobrew -l meson
213

$ ./gobrew -l swig
32

$ ./gobrew -l llvm
51
```

### show all build dependencies and their count

```sh
$ ./gobrew -b
All Build Dependencies Count:  320
[ghostscript openjdk@8 libsigc++@2 ucl m4 pcre scala gengetopt automake ghc@9.6 tl-expected xmake libtool docbook-xsl perl cweb lua pyinvoke cmake swig coreutils imagemagick gperf spdlog gnu-tar libevent sphinx-doc gputils gsettings-desktop-schemas tcl-tk unifdef gi-docgen graphviz go-md2man nasm bind freetds json-glib xbyak rustup-init counterfeiter vulkan-loader cpptest cargo-bundle libelf mkfontscale cargo-c jinja2-cli tradcpp pnpm ocaml-findlib gtkmm tclap libx11 libgpg-error cairomm@1.14 pipenv llvm go@1.17 python@3.11 node gradle@6 libpq linux-headers@5.15 vim sip nlohmann-json protobuf dune nkf openssl@3 doctest ghc@8.10 gtk-doc yarn hidapi sbt mage dos2unix asciidoctor hpack gmp flatbuffers t1utils keystone ronn ocaml pandoc go@1.19 lzip miniupnpc xorgproto opam mvfst docbook2x texinfo cc65 libxkbfile gatsby-cli mdds protobuf@21 go cppunit mysql@5.7 linux-headers@4.4 glide python-setuptools ghc@8.6 ghc@9.2 gox lpeg node@14 ecl util-linux glm asciidoc go-bindata spirv-llvm-translator gpp boost gobject-introspection boost-build opencl-headers vala dub libetonyek valijson python@3.12 maven sdcc gnutls doxygen gnustep-make autoconf@2.13 socat breezy gcc mandoc jam llvm@16 luarocks ocaml@4 librsvg cereal mercurial erlang gzip intltool pillow numpy eigen ant glib extra-cmake-modules binutils texi2html xz bmake mk-configure cabal-install libatomic_ops help2man libuv vulkan-volk ghc@9.4 bash groff meson nginx nim google-sparsehash qt python@3.9 googletest autoconf-archive opus freetype ki18n httpd lndir docbook po4a gawk node@18 ruby tmux tlx libosmium gtk4 yasm qt@5 check wabt xmltoman bsdmake apr grep pyqt-builder pod2man ispc rsync cython re2c cpanminus pybind11 cscope cxxopts util-macros xcb-proto cli11 osinfo-db-tools ford apr-util fpc ifacemaker font-util yelp-tools python@3.10 libgcrypt ghc bazelisk go@1.20 pkg-config bat libwpg libxt openjdk@11 gnome-common pypy spice-protocol python-lxml byacc autoconf@2.69 scons texlive msgpack-cxx halibut uthash dtools rebar3 nettle libxslt openjdk flex pygobject3 glktermw lit helm gettext itstool shfmt zig antlr scdoc haskell-stack fmt molten-vk autoconf dmd clojure asio gnu-getopt crystal libarchive rapidjson xmlto gradle smake quickjs docutils w3m bison cmocka gnu-sed libconfig cpptoml mingw-w64 dotnet xa pyyaml desktop-file-utils make hevea lowdown cunit foma pangomm@2.46 go@1.18 gnupg poetry chafa emacs netsurf-buildsystem sassc rpcgen fontforge ocamlbuild ragel argp-standalone xtrans rust rustfmt ldc just glibmm@2.66 vulkan-headers libsodium mlton buildapp pcre2 libscrypt erlang@25 mockery mcpp repo mlkit ninja go@1.21 ccache luajit]

$ ./gobrew --build-dep
All Build Dependencies Count:  320
[xmake grep ronn gobject-introspection ocaml font-util mkfontscale lpeg ucl vulkan-volk yarn glib ragel libconfig mandoc xtrans bat libosmium automake flex xz fmt clojure cpanminus vulkan-loader docutils protobuf@21 jinja2-cli doxygen extra-cmake-modules gatsby-cli httpd halibut rustup-init ruby cargo-c crystal miniupnpc dtools dub node@14 gtkmm msgpack-cxx socat boost rustfmt cweb foma binutils wabt googletest libtool libxslt hpack asio meson poetry ifacemaker python-setuptools desktop-file-utils cxxopts glktermw byacc xbyak openjdk@11 uthash bind go-bindata graphviz argp-standalone protobuf eigen autoconf@2.69 gnu-getopt libgpg-error libpq ocaml@4 mockery libuv flatbuffers ghc@8.10 pangomm@2.46 coreutils emacs unifdef lndir cmake bsdmake dotnet librsvg libevent python@3.9 nettle mingw-w64 cli11 perl gmp freetds pybind11 openssl@3 rpcgen docbook xmltoman libgcrypt util-linux go@1.21 pyinvoke ocaml-findlib mysql@5.7 apr pod2man pandoc erlang@25 gnustep-make spirv-llvm-translator libxkbfile itstool rebar3 libetonyek libsigc++@2 pyqt-builder autoconf boost-build tcl-tk apr-util cscope autoconf-archive shfmt openjdk hidapi mvfst openjdk@8 tclap maven pcre2 breezy asciidoctor libwpg jam keystone valijson scdoc hevea mlton groff libelf bash opencl-headers osinfo-db-tools sip dos2unix gnutls xa ghc quickjs erlang chafa gperf vim antlr cunit pkg-config yasm lowdown netsurf-buildsystem go@1.18 pypy re2c lua gzip gnupg ghostscript nim docbook-xsl xmlto ldc counterfeiter yelp-tools xcb-proto bison mcpp pipenv rust python@3.11 imagemagick ispc tmux cairomm@1.14 gettext glm gcc pygobject3 ecl go@1.17 libx11 pcre gnu-tar libxt node scala util-macros python@3.10 node@18 freetype lzip swig vala m4 sassc ocamlbuild gsettings-desktop-schemas pnpm go-md2man gawk gnome-common go llvm pillow smake nkf ccache linux-headers@4.4 buildapp repo gtk4 cabal-install sphinx-doc xorgproto go@1.19 tlx ghc@9.4 ford ki18n doctest cpptoml po4a luarocks opus gi-docgen ghc@9.2 spice-protocol dune tl-expected gradle@6 tradcpp intltool mk-configure rsync libsodium glibmm@2.66 llvm@16 ant help2man qt check ghc@9.6 sdcc cpptest luajit json-glib gradle qt@5 texi2html docbook2x bmake fontforge autoconf@2.13 python-lxml vulkan-headers spdlog cc65 helm haskell-stack dmd just texinfo texlive cython gpp molten-vk cereal rapidjson nasm cmocka mlkit mage bazelisk gputils gtk-doc libscrypt asciidoc fpc gox nlohmann-json cppunit w3m mdds lit t1utils scons gengetopt google-sparsehash python@3.12 libatomic_ops gnu-sed zig sbt cargo-bundle pyyaml make ghc@8.6 numpy mercurial go@1.20 linux-headers@5.15 nginx glide ninja libarchive opam]
```

## tasks

- [x] get all Homebrew Core formulas
- [x] save core_formulas as a file onto the disk
- [x] get each package JSON file
- [x] get count of packages which are written/built in Go language
- [x] ability to set the language or build system
- [x] include **dependencies** in calculation
- [x] include **build dependencies** in calculation
- [x] include **test dependencies** in calculation
- [x] include **recommended dependencies** in calculation
- [x] include **optional dependencies** in calculation
- [x] count all versions of the language by default ([commit](https://github.com/abanoubha/gobrew/commit/7de9e76c03401ce70568417db550eda590bff919))
- [x] re-download Homebrew/Core formulae index JSON file if the local one is older than 7 days ([commit](https://github.com/abanoubha/gobrew/commit/2a9713b90dd319203ec7692df81fb6c8e5759277))
