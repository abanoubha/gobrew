# Documented development processes

## install gobrew

- install using Go

```sh
go install github.com/abanoubha/gobrew

# or

go install github.com/abanoubha/gobrew@latest
```

## build executables of all operating systems / platforms

Just use the `build-all.sh` POSIX compliant bash script, like this:

```sh
sh scripts/build-all.sh v260205
```

`v260205` is the version of release you're building/compiling. The output executable files will be `gobrew-OS-ARCH-VERSION`.

Or you can use direct commands like this:

```sh
# linux 64 bit (current os)
go build -o gobrew-linux-x64 .
# linux 64 bit (if not working on Linux distro)
GOOS=linux GOARCH=amd64 go build -o gobrew-linux-x64 .

# windows 64 bit
GOOS=windows GOARCH=amd64 go build -o gobrew-windows-x64.exe .

# macOS M-series
GOOS=darwin GOARCH=arm64 go build -o gobrew-macos-apple-silicon .
# macOS intel 64 bit
GOOS=darwin GOARCH=amd64 go build -o gobrew-macos-x64 .
```

## list all TODOs

use `grep` to list all TODO comments:

```sh
grep -rni "TODO" --include="*.go" --exclude-dir=.git 2> /dev/null

# or just use
grep -rni "todo"
```

or use `rg` to list all TODOs:

```sh
rg --type-add 'go:*.go' "TODO|FIXME|todo|fixme|fix" --glob '!.git/'
```
