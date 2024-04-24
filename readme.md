# gobrew

Count all programs written/built in Go/Rust/Python/.. and distributed via Homebrew.

## why I build this?

I was curious. Are people *actually* using Go more than Rust? Are there more software written with cmake or meson or ninja .. ?

Too many curious questions I wanted to answer them using *realworld statistics*.

## commands

```sh
# get deps, build, run
# will show the count of packages written in Go in Homebrew Core (by default)
$ go mod tidy && go build -o gobrew main.go && ./gobrew
957

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l go
957

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l rust
524

# get the packages built with "ninja" build system
$ go mod tidy && go build -o gobrew main.go && ./gobrew -l ninja
249

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l cython
8

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l gcc
6

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l cmake
1005

$ go mod tidy && go build -o gobrew main.go && ./gobrew --lang meson
213

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l ruby
2

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l swig
28

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l lua
6

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l python@3.12
122

$ ./gobrew
957

$ ./gobrew -l rust
524

$ gobrew
957

$ gobrew -l rust
524
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
