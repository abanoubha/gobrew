# gobrew

[![gobrew](https://snapcraft.io/gobrew/badge.svg)](https://snapcraft.io/gobrew)

Count all programs written/built in Go/Rust/Python/.. and distributed via Homebrew.

[![asciicast](https://asciinema.org/a/674093.svg)](https://asciinema.org/a/674093)

Check out the [change log and roadmap](CHANGELOG.md)!

## why I built this?

I was curious. Are people _actually_ using Go more than Rust? Are there more software written with cmake or meson or ninja .. ?

Too many curious questions I wanted to answer them using _realworld statistics_.

I wrote a post about my findings regarding [Go vs Rust usage in Homebrew Core formulae on my website](https://abanoubhanna.com/posts/go-vs-rust-use-production/).

## install gobrew

- install via the snap store

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/gobrew)

- install using Go

```sh
go install github.com/abanoubha/gobrew@latest
```

- install via piping a script into bash

```sh
curl -fsSL https://raw.githubusercontent.com/abanoubha/gobrew/main/scripts/install.sh | sh
```

## build CLI app from source

```sh
# get deps, build the "gobrew" executable
$ go mod tidy && go build -o gobrew .
```

## uninstall gobrew

- Remove gobrew executable using the script:

```sh
curl -fsSL https://raw.githubusercontent.com/abanoubha/gobrew/main/scripts/uninstall.sh | sh
```

If you want to specify the installation directory/folder, add `INSTALL_DIR="/usr/local/bin"` to before bash like this:

```sh
curl -fsSL https://raw.githubusercontent.com/abanoubha/gobrew/main/scripts/uninstall.sh | INSTALL_DIR="/usr/local/bin" sh
```

If you have renamed the executable/binary, specify it with `INSTALL_NAME="gobrew"` like this:

```sh
curl -fsSL https://raw.githubusercontent.com/abanoubha/gobrew/main/scripts/uninstall.sh | INSTALL_NAME="gobrew" INSTALL_DIR="/usr/local/bin" sh
```

- if you install gobrew via _snap_:

```sh
sudo snap remove gobrew
```

- if you install gobrew via Go toolchain:

```sh
go clean -i github.com/abanoubha/gobrew
```

## count all packages that use a specific language

- will show the count of packages written in Go in Homebrew Core (by default)

```sh
$ gobrew
No language nor build system nor library is specified. Counting packages built in Go (by default):
974
```

- count all programs written in Go language

```sh
$ gobrew -l go
974
```

or

```sh
$ gobrew --lang go
974
```

- count all programs written in Rust

```sh
$ gobrew -l rust
524
```

- count all programs written in Zig programming languages

```sh
$ gobrew -l zig
6
```

## specific version vs. all versions of the specified language

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

## count all packages that use a specific build system or library

- count the packages built with "ninja" build system

```sh
$ gobrew -l ninja
253
```

- count all packages/programs that use/rely-on "gcc":

```sh
$ gobrew -l gcc
75
```

- count all programs that use cmake

```sh
$ gobrew -l cmake
1011
```

- count all apps that depend on meson as a builder

```sh
$ gobrew -l meson
213
```

- count all programs that use LLVM

```sh
$ gobrew -l llvm
51
```

## show all build dependencies and their count

```sh
$ gobrew -b
All Build Dependencies Count:  320
[ghostscript openjdk@8 libsigc++@2 ucl m4 pcre scala gengetopt automake ghc@9.6 tl-expected xmake libtool docbook-xsl perl cweb lua pyinvoke cmake swig coreutils imagemagick gperf spdlog gnu-tar libevent sphinx-doc gputils ... trucated]

$ gobrew --build-dep
All Build Dependencies Count:  320
[xmake grep ronn gobject-introspection ocaml font-util mkfontscale lpeg ucl vulkan-volk yarn glib ragel libconfig mandoc xtrans bat libosmium automake flex xz fmt clojure cpanminus vulkan-loader docutils protobuf@21 jinja2-cli doxygen ... truncated]
```

## show all languages/libs and the count of their packages

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
# truncated for brevity ...
```

## show all packages which depend on a certain lib/lang

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
