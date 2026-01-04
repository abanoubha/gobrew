# gobrew

[![gobrew](https://snapcraft.io/gobrew/badge.svg)](https://snapcraft.io/gobrew)

Count all programs written/built in Go/Rust/Python/.. and distributed via Homebrew.

[![asciicast](https://asciinema.org/a/674093.svg)](https://asciinema.org/a/674093)

Check out the [change log and roadmap](CHANGELOG.md)!

## why I built this?

I was curious. Are people _actually_ using Go more than Rust? Are there more software written with cmake or meson or ninja .. ?

Too many curious questions I wanted to answer them using _realworld statistics_.

I wrote a post about my findings regarding [Go vs Rust usage in Homebrew Core formulae on my website](https://abanoubhanna.com/posts/go-vs-rust-use-production/).

## commands

### install gobrew

- install using Go

```sh
go install github.com/abanoubha/gobrew
```

- install via the snap store

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/gobrew)

### build CLI app from source

```sh
# get deps, build the "gobrew" executable
$ go mod tidy && go build -o gobrew .
```

### count all packages that use a specific language

```sh
# will show the count of packages written in Go in Homebrew Core (by default)
$ gobrew
No language nor build system nor library is specified. Counting packages built in Go (by default):
974

$ gobrew -l go
974

$ gobrew --lang go
974

$ gobrew -l rust
524

$ gobrew -l cython
12

$ gobrew -l ruby
27

$ gobrew -l lua
51

$ gobrew -l zig
6
```

#### all versions of the specified language

- all versions of python

```sh
$ gobrew -l python
771
```

- specific version of python

```sh
$ gobrew -l python@3.12
718

$ gobrew -l python@3.11
67

$ gobrew -l python@3.10
10

$ gobrew -l python@3.9
6

$ gobrew -l python@3.8
2
```

### count all packages that use a specific build system or library

```sh
# get the packages built with "ninja" build system
$ gobrew -l ninja
253

$ gobrew -l gcc
75

$ gobrew -l cmake
1011

$ gobrew -l meson
213

$ gobrew -l swig
32

$ gobrew -l llvm
51
```

### show all build dependencies and their count

```sh
$ gobrew -b
All Build Dependencies Count:  320
[ghostscript openjdk@8 libsigc++@2 ucl m4 pcre scala gengetopt automake ghc@9.6 tl-expected xmake libtool docbook-xsl perl cweb lua pyinvoke cmake swig coreutils imagemagick gperf spdlog gnu-tar libevent sphinx-doc gputils gsettings-desktop-schemas tcl-tk unifdef gi-docgen graphviz go-md2man nasm bind freetds json-glib xbyak rustup-init counterfeiter vulkan-loader cpptest cargo-bundle libelf mkfontscale cargo-c jinja2-cli tradcpp pnpm ocaml-findlib gtkmm tclap libx11 libgpg-error cairomm@1.14 pipenv llvm go@1.17 python@3.11 node gradle@6 libpq linux-headers@5.15 vim sip nlohmann-json protobuf dune nkf openssl@3 doctest ghc@8.10 gtk-doc yarn hidapi sbt mage dos2unix asciidoctor hpack gmp flatbuffers t1utils keystone ronn ocaml pandoc go@1.19 lzip miniupnpc xorgproto opam mvfst docbook2x texinfo cc65 libxkbfile gatsby-cli mdds protobuf@21 go cppunit mysql@5.7 linux-headers@4.4 glide python-setuptools ghc@8.6 ghc@9.2 gox lpeg node@14 ecl util-linux glm asciidoc go-bindata spirv-llvm-translator gpp boost gobject-introspection boost-build opencl-headers vala dub libetonyek valijson python@3.12 maven sdcc gnutls doxygen gnustep-make autoconf@2.13 socat breezy gcc mandoc jam llvm@16 luarocks ocaml@4 librsvg cereal mercurial erlang gzip intltool pillow numpy eigen ant glib extra-cmake-modules binutils texi2html xz bmake mk-configure cabal-install libatomic_ops help2man libuv vulkan-volk ghc@9.4 bash groff meson nginx nim google-sparsehash qt python@3.9 googletest autoconf-archive opus freetype ki18n httpd lndir docbook po4a gawk node@18 ruby tmux tlx libosmium gtk4 yasm qt@5 check wabt xmltoman bsdmake apr grep pyqt-builder pod2man ispc rsync cython re2c cpanminus pybind11 cscope cxxopts util-macros xcb-proto cli11 osinfo-db-tools ford apr-util fpc ifacemaker font-util yelp-tools python@3.10 libgcrypt ghc bazelisk go@1.20 pkg-config bat libwpg libxt openjdk@11 gnome-common pypy spice-protocol python-lxml byacc autoconf@2.69 scons texlive msgpack-cxx halibut uthash dtools rebar3 nettle libxslt openjdk flex pygobject3 glktermw lit helm gettext itstool shfmt zig antlr scdoc haskell-stack fmt molten-vk autoconf dmd clojure asio gnu-getopt crystal libarchive rapidjson xmlto gradle smake quickjs docutils w3m bison cmocka gnu-sed libconfig cpptoml mingw-w64 dotnet xa pyyaml desktop-file-utils make hevea lowdown cunit foma pangomm@2.46 go@1.18 gnupg poetry chafa emacs netsurf-buildsystem sassc rpcgen fontforge ocamlbuild ragel argp-standalone xtrans rust rustfmt ldc just glibmm@2.66 vulkan-headers libsodium mlton buildapp pcre2 libscrypt erlang@25 mockery mcpp repo mlkit ninja go@1.21 ccache luajit]

$ gobrew --build-dep
All Build Dependencies Count:  320
[xmake grep ronn gobject-introspection ocaml font-util mkfontscale lpeg ucl vulkan-volk yarn glib ragel libconfig mandoc xtrans bat libosmium automake flex xz fmt clojure cpanminus vulkan-loader docutils protobuf@21 jinja2-cli doxygen extra-cmake-modules gatsby-cli httpd halibut rustup-init ruby cargo-c crystal miniupnpc dtools dub node@14 gtkmm msgpack-cxx socat boost rustfmt cweb foma binutils wabt googletest libtool libxslt hpack asio meson poetry ifacemaker python-setuptools desktop-file-utils cxxopts glktermw byacc xbyak openjdk@11 uthash bind go-bindata graphviz argp-standalone protobuf eigen autoconf@2.69 gnu-getopt libgpg-error libpq ocaml@4 mockery libuv flatbuffers ghc@8.10 pangomm@2.46 coreutils emacs unifdef lndir cmake bsdmake dotnet librsvg libevent python@3.9 nettle mingw-w64 cli11 perl gmp freetds pybind11 openssl@3 rpcgen docbook xmltoman libgcrypt util-linux go@1.21 pyinvoke ocaml-findlib mysql@5.7 apr pod2man pandoc erlang@25 gnustep-make spirv-llvm-translator libxkbfile itstool rebar3 libetonyek libsigc++@2 pyqt-builder autoconf boost-build tcl-tk apr-util cscope autoconf-archive shfmt openjdk hidapi mvfst openjdk@8 tclap maven pcre2 breezy asciidoctor libwpg jam keystone valijson scdoc hevea mlton groff libelf bash opencl-headers osinfo-db-tools sip dos2unix gnutls xa ghc quickjs erlang chafa gperf vim antlr cunit pkg-config yasm lowdown netsurf-buildsystem go@1.18 pypy re2c lua gzip gnupg ghostscript nim docbook-xsl xmlto ldc counterfeiter yelp-tools xcb-proto bison mcpp pipenv rust python@3.11 imagemagick ispc tmux cairomm@1.14 gettext glm gcc pygobject3 ecl go@1.17 libx11 pcre gnu-tar libxt node scala util-macros python@3.10 node@18 freetype lzip swig vala m4 sassc ocamlbuild gsettings-desktop-schemas pnpm go-md2man gawk gnome-common go llvm pillow smake nkf ccache linux-headers@4.4 buildapp repo gtk4 cabal-install sphinx-doc xorgproto go@1.19 tlx ghc@9.4 ford ki18n doctest cpptoml po4a luarocks opus gi-docgen ghc@9.2 spice-protocol dune tl-expected gradle@6 tradcpp intltool mk-configure rsync libsodium glibmm@2.66 llvm@16 ant help2man qt check ghc@9.6 sdcc cpptest luajit json-glib gradle qt@5 texi2html docbook2x bmake fontforge autoconf@2.13 python-lxml vulkan-headers spdlog cc65 helm haskell-stack dmd just texinfo texlive cython gpp molten-vk cereal rapidjson nasm cmocka mlkit mage bazelisk gputils gtk-doc libscrypt asciidoc fpc gox nlohmann-json cppunit w3m mdds lit t1utils scons gengetopt google-sparsehash python@3.12 libatomic_ops gnu-sed zig sbt cargo-bundle pyyaml make ghc@8.6 numpy mercurial go@1.20 linux-headers@5.15 nginx glide ninja libarchive opam]
```

### show all languages/libs and the count of their packages

```sh
$ gobrew -s
count of all languages/libraries/frameworks:  1489
pkg-config : 1554
cmake : 1074
go : 1008
python@3.12 : 791
rust : 553
openssl@3 : 519
autoconf : 459
automake : 443
libtool : 369
gettext : 280
ninja : 256
openjdk : 230
meson : 216
glib : 214
node : 198
libpng : 181
boost : 174
certifi : 174
libyaml : 161
jpeg-turbo : 135
readline : 128
gmp : 115
zstd : 114
freetype : 113
gobject-introspection : 96
pcre2 : 95
xz : 93
python@3.11 : 89
sdl2 : 89
icu4c : 74
gcc : 73
libusb : 71
gtk+3 : 71
libtiff : 70
libx11 : 69
bison : 68
cryptography : 66
cairo : 65
libvorbis : 61
sqlite : 61
gnutls : 61
lz4 : 58
libevent : 56
doxygen : 55
pcre : 54
python-setuptools : 53
ncurses : 52
vala : 49
lua : 49
texinfo : 49
qt : 48
# truncated for brevity ...
```

### show all packages which depend on a certain lib/lang

```sh
$ gobrew -d zig

 zigmod :
   Package manager for the Zig programming language

 zls :
   Language Server for Zig

 cargo-zigbuild :
   Compile Cargo project with zig as linker

 fastfec :
   Extremely fast FEC filing parser written in C

 ncdu :
   NCurses Disk Usage

 zf :
   Command-line fuzzy finder that prioritizes matches on filenames
```
