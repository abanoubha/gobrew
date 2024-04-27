# gobrew

Count all programs written/built in Go/Rust/Python/.. and distributed via Homebrew.

## why I build this?

I was curious. Are people *actually* using Go more than Rust? Are there more software written with cmake or meson or ninja .. ?

Too many curious questions I wanted to answer them using *realworld statistics*.

## commands

### build CLI app from source

```sh
# get deps, build the "gobrew" executable
$ go mod tidy && go build -o gobrew main.go
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
