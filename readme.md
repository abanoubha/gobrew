# gobrew

Count all programs written/built in Go/Rust/Python/.. and distributed via Homebrew.

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

### show all languages/libs and the count of their packages

```sh
$ ./gobrew -s
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
cabal-install : 47
libogg : 47
fontconfig : 47
numpy : 47
mpfr : 44
pango : 44
protobuf : 44
libgcrypt : 43
coreutils : 42
llvm : 42
libpq : 42
openblas : 41
xorgproto : 40
intltool : 40
ffmpeg : 39
fftw : 39
openjdk@11 : 39
libsodium : 38
libgit2 : 37
libarchive : 36
eigen : 36
fmt : 36
flac : 34
libxext : 34
webp : 34
harfbuzz : 33
sdl12-compat : 33
libsndfile : 33
swig : 32
little-cms2 : 32
ruby : 31
pandoc : 31
libxml2 : 31
ocaml : 30
gdk-pixbuf : 30
adwaita-icon-theme : 30
sphinx-doc : 30
libuv : 29
gnu-sed : 29
abseil : 28
docbook-xsl : 28
zeromq : 28
rustup-init : 28
qt@5 : 27
graphviz : 27
curl : 27
openjpeg : 26
libomp : 26
json-glib : 26
imagemagick : 26
open-mpi : 26
maven : 26
libgpg-error : 25
json-c : 25
perl : 25
ghc : 25
libzip : 24
luajit : 24
bash : 23
giflib : 23
gtk+ : 23
hdf5 : 23
ghostscript : 23
pillow : 23
snappy : 22
jansson : 22
util-macros : 22
help2man : 21
libfuse@2 : 21
gpgme : 20
libmpc : 20
unixodbc : 20
pygobject3 : 20
tbb : 20
php : 20
bdw-gc : 20
zlib : 20
gd : 20
xmlto : 19
libxt : 19
gflags : 19
poppler : 19
yarn : 19
opus : 19
lame : 19
lzo : 19
c-ares : 18
popt : 18
flex : 18
tcl-tk : 18
sdl2_image : 18
itstool : 18
glew : 18
glog : 18
libsamplerate : 18
libidn2 : 18
libxcb : 18
librsvg : 17
gnupg : 17
brotli : 17
sdl2_mixer : 17
libnghttp2 : 17
gsl : 16
libsoup : 16
libmagic : 16
gtk4 : 16
libev : 16
openjdk@17 : 16
at-spi2-core : 16
asciidoc : 16
libusb-compat : 16
gawk : 15
ant : 15
libssh2 : 15
kubernetes-cli : 15
apr : 15
openexr : 15
nettle : 15
proj : 15
krb5 : 14
jsoncpp : 14
openjdk@8 : 14
isl : 14
ffmpeg@6 : 14
asciidoctor : 14
nlohmann-json : 14
berkeley-db@5 : 14
portaudio : 14
erlang : 14
libxmu : 14
libao : 14
mysql-client : 13
wxwidgets : 13
openjdk@21 : 13
protobuf@21 : 13
mad : 13
groonga : 13
emacs : 13
ldc : 13
nasm : 13
glm : 13
vulkan-headers : 13
xerces-c : 13
minizip : 13
libaec : 13
openssl@1.1 : 13
googletest : 13
jq : 13
gradle : 13
netcdf : 13
fluid-synth : 13
libnet : 12
docbook : 12
gdbm : 12
ghc@9.4 : 12
librevenge : 12
apr-util : 12
speex : 12
libice : 12
hwloc : 12
hicolor-icon-theme : 12
theora : 12
re2 : 12
protobuf-c : 12
sdl2_ttf : 11
ghc@9.6 : 11
python@3.10 : 11
libsoxr : 11
gtk-doc : 11
cython : 11
pybind11 : 11
glfw : 11
imath : 11
scdoc : 11
jpeg-xl : 11
postgresql@14 : 11
oniguruma : 11
ca-certificates : 11
libssh : 11
jack : 11
libexif : 11
rbenv : 11
httpd : 11
cfitsio : 11
suite-sparse : 11
libsm : 10
dbus : 10
hidapi : 10
tesseract : 10
terminal-notifier : 10
boost-python3 : 10
taglib : 10
opam : 10
pixman : 10
groff : 10
libfuse : 10
geos : 10
vulkan-loader : 10
libheif : 9
tmux : 9
extra-cmake-modules : 9
py3cairo : 9
yasm : 9
tinyxml2 : 9
haskell-stack : 9
docutils : 9
libsigc++@2 : 9
pugixml : 9
rapidjson : 9
libpcap : 9
jasper : 9
lmdb : 9
libadwaita : 9
gnuplot : 9
mesa : 9
libmaxminddb : 9
libunistring : 9
libfido2 : 9
opusfile : 9
libmicrohttpd : 9
expat : 9
asio : 9
ocaml-findlib : 9
mpg123 : 9
m4 : 8
s-lang : 8
libvpx : 8
libassuan : 8
argon2 : 8
glibmm@2.66 : 8
boost@1.76 : 8
grpc : 8
z3 : 8
systemd : 8
gdal : 8
libxft : 8
libxrandr : 8
fribidi : 8
libidn : 8
gstreamer : 8
spdlog : 8
folly : 8
libxinerama : 8
capstone : 8
libxfixes : 8
libplist : 8
qhull : 8
gspell : 8
openldap : 8
desktop-file-utils : 8
mage : 7
libgsf : 7
pod2man : 7
libxpm : 7
libmpdclient : 7
libspatialite : 7
gnu-getopt : 7
xxhash : 7
x264 : 7
liblo : 7
aspell : 7
xorg-server : 7
utf8proc : 7
libxrender : 7
libid3tag : 7
scons : 7
libass : 7
metis : 7
elfutils : 7
dotnet : 7
xtrans : 7
double-conversion : 7
autoconf-archive : 7
miniupnpc : 7
netpbm : 7
molten-vk : 7
fltk : 7
libftdi : 7
node@18 : 6
libgee : 6
mbedtls@2 : 6
htslib : 6
sdl2_net : 6
go-md2man : 6
enchant : 6
wget : 6
mecab : 6
aom : 6
djvulibre : 6
x265 : 6
libtasn1 : 6
libxslt : 6
libxi : 6
hiredis : 6
zimg : 6
yaml-cpp : 6
p7zip : 6
libconfig : 6
opencv : 6
libdnet : 6
libwpd : 6
cjson : 6
argp-standalone : 6
freeimage : 6
cffi : 6
tokyo-cabinet : 6
libmnl : 6
freetds : 6
pari : 6
faad2 : 6
sbcl : 6
helm : 6
make : 6
frei0r : 6
opencl-headers : 6
util-linux : 6
confuse : 6
ghc@9.2 : 6
speexdsp : 6
libmodplug : 6
check : 6
scipy : 6
rtmpdump : 6
zig : 6
leveldb : 5
raptor : 5
mpdecimal : 5
libnfc : 5
dav1d : 5
macos-term-size : 5
pyqt@5 : 5
wavpack : 5
rebar3 : 5
coinutils : 5
unbound : 5
grep : 5
cairomm@1.14 : 5
libcap : 5
libbluray : 5
libtommath : 5
srt : 5
lzlib : 5
docker : 5
luarocks : 5
fizz : 5
octomap : 5
bzip2 : 5
sdl_image : 5
libepoxy : 5
jemalloc : 5
neon : 5
mono : 5
rubberband : 5
sdl_ttf : 5
libraw : 5
opencore-amr : 5
assimp : 5
gi-docgen : 5
crystal : 5
tidy-html5 : 5
xvid : 5
fzf : 5
bazelisk : 5
pulseaudio : 5
libxmp : 5
xapian : 5
librist : 5
gperf : 5
glslang : 5
vapoursynth : 5
git : 5
guile : 5
bsdmake : 5
llvm@15 : 5
yajl : 5
gperftools : 5
neovim : 5
rav1e : 5
exiv2 : 4
dune : 4
physfs : 4
cargo-c : 4
libvmaf : 4
purescript : 4
lzip : 4
vips : 4
orc : 4
cgal : 4
gsettings-desktop-schemas : 4
exiftool : 4
osi : 4
libshout : 4
pkcs11-helper : 4
sdl_mixer : 4
dialog : 4
lv2 : 4
libslirp : 4
flann : 4
pyenv : 4
llvm@17 : 4
linux-headers@5.15 : 4
pinentry : 4
xcb-util : 4
libvidstab : 4
sbt : 4
go-bindata : 4
libdvdread : 4
uriparser : 4
soapysdr : 4
libxp : 4
task : 4
bmake : 4
glib-networking : 4
ceres-solver : 4
libxaw : 4
libffi : 4
uthash : 4
go@1.21 : 4
librdkafka : 4
vte3 : 4
cppunit : 4
vtk : 4
libsoup@2 : 4
libsecret : 4
libxau : 4
msgpack : 4
apache-arrow : 4
uchardet : 4
scalapack : 4
tree-sitter : 4
libseccomp : 4
zsh : 4
serd : 4
rlwrap : 4
ghc@8.10 : 4
libavif : 4
gtksourceview4 : 4
mercurial : 4
graphene : 4
game-music-emu : 4
freexl : 4
socat : 4
w3m : 4
```

### show all packages which depend on a certain lib/lang

```sh
$ ./gobrew -d zig

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
- [x] show all languages and count of their packages depend on them
- [x] show all packages depends on specific language
