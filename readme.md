# gobrew

Count all programs written in Go and distributed via Homebrew.

## commands

```sh
# get deps, build, run
# will show the count of packages written in Go in Homebrew Core (by default)
$ go mod tidy && go build -o gobrew main.go && ./gobrew
938

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l go
938

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l rust
522

# get the packages built with "ninja" build system
$ go mod tidy && go build -o gobrew main.go && ./gobrew -l ninja
249

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l cython
8

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l gcc
6

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l cmake
1005

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l ruby
2

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l swig
28

$ go mod tidy && go build -o gobrew main.go && ./gobrew -l lua
6
```

## tasks

- [x] get all Homebrew Core formulas
- [x] save core_formulas as a file onto the disk
- [x] get each package JSON file
- [x] get count of packages which are written/built in Go language
- [x] ability to set the language or build system
