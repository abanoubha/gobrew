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

## executable/binary size reduction

For "normal" build command which is `GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$OUT_DIR/$OUTPUT_NAME" "$PACKAGE_PATH"` in [build-all.sh](./scripts/build-all.sh), binary sizes are:

```sh
$ eza -alb dist/
.rwxr-xr-x 9.2Mi aba  6 Mar 16:29 gobrew-freebsd-amd64-v260306
.rwxr-xr-x 8.6Mi aba  6 Mar 16:29 gobrew-freebsd-arm64-v260306
.rwxr-xr-x 8.5Mi aba  6 Mar 16:29 gobrew-freebsd-riscv64-v260306
.rwxr-xr-x 8.6Mi aba  6 Mar 16:27 gobrew-linux-arm64-v260306
.rwxr-xr-x 9.4Mi aba  6 Mar 16:27 gobrew-linux-x64-v260306
.rwxr-xr-x 8.9Mi aba  6 Mar 16:29 gobrew-macos-apple-silicon-arm64-v260306
.rwxr-xr-x 9.5Mi aba  6 Mar 16:29 gobrew-macos-intel-x64-v260306
.rwxr-xr-x 9.2Mi aba  6 Mar 16:30 gobrew-netbsd-amd64-v260306
.rwxr-xr-x 8.5Mi aba  6 Mar 16:30 gobrew-netbsd-arm64-v260306
.rwxr-xr-x 9.2Mi aba  6 Mar 16:30 gobrew-openbsd-amd64-v260306
.rwxr-xr-x 8.5Mi aba  6 Mar 16:30 gobrew-openbsd-arm64-v260306
.rwxr-xr-x 8.3Mi aba  6 Feb 10:09 gobrew-windows-arm64-v260206.exe
.rwxr-xr-x 8.6Mi aba  6 Mar 16:28 gobrew-windows-arm64-v260306.exe
.rwxr-xr-x 9.1Mi aba  6 Feb 10:09 gobrew-windows-x64-v260206.exe
.rwxr-xr-x 9.4Mi aba  6 Mar 16:28 gobrew-windows-x64-v260306.exe
```

But after adding `-trimpath -ldflags="-s -w -buildid="` to strip debug info and unwanted data from the binary - the command becomes `GOOS="$GOOS" GOARCH="$GOARCH" go build -trimpath -ldflags="-s -w -buildid=" -o "$OUT_DIR/$OUTPUT_NAME" "$PACKAGE_PATH"` in [build-all.sh](./scripts/build-all.sh), binary sizes became:

```sh
$ eza -alb dist/
.rwxr-xr-x 6.4Mi aba  6 Mar 17:58 gobrew-freebsd-amd64-v260306
.rwxr-xr-x 6.0Mi aba  6 Mar 17:58 gobrew-freebsd-arm64-v260306
.rwxr-xr-x 5.8Mi aba  6 Mar 17:59 gobrew-freebsd-riscv64-v260306
.rwxr-xr-x 6.0Mi aba  6 Mar 17:57 gobrew-linux-arm64-v260306
.rwxr-xr-x 6.5Mi aba  6 Mar 17:57 gobrew-linux-x64-v260306
.rwxr-xr-x 6.1Mi aba  6 Mar 17:58 gobrew-macos-apple-silicon-arm64-v260306
.rwxr-xr-x 6.6Mi aba  6 Mar 17:58 gobrew-macos-intel-x64-v260306
.rwxr-xr-x 6.4Mi aba  6 Mar 17:59 gobrew-netbsd-amd64-v260306
.rwxr-xr-x 5.9Mi aba  6 Mar 17:59 gobrew-netbsd-arm64-v260306
.rwxr-xr-x 6.4Mi aba  6 Mar 17:59 gobrew-openbsd-amd64-v260306
.rwxr-xr-x 5.9Mi aba  6 Mar 18:00 gobrew-openbsd-arm64-v260306
.rwxr-xr-x 8.3Mi aba  6 Feb 10:09 gobrew-windows-arm64-v260206.exe
.rwxr-xr-x 6.0Mi aba  6 Mar 17:57 gobrew-windows-arm64-v260306.exe
.rwxr-xr-x 9.1Mi aba  6 Feb 10:09 gobrew-windows-x64-v260206.exe
.rwxr-xr-x 6.7Mi aba  6 Mar 17:57 gobrew-windows-x64-v260306.exe
```

Here is the difference:

```plain
.rwxr-xr-x 9.2Mi aba  6 Mar 16:29 gobrew-freebsd-amd64-v260306
.rwxr-xr-x 6.4Mi aba  6 Mar 17:58 gobrew-freebsd-amd64-v260306
---------- 2.8Mi ---------------------------------------------

.rwxr-xr-x 8.6Mi aba  6 Mar 16:29 gobrew-freebsd-arm64-v260306
.rwxr-xr-x 6.0Mi aba  6 Mar 17:58 gobrew-freebsd-arm64-v260306
---------- 2.6Mi ---------------------------------------------

.rwxr-xr-x 8.5Mi aba  6 Mar 16:29 gobrew-freebsd-riscv64-v260306
.rwxr-xr-x 5.8Mi aba  6 Mar 17:59 gobrew-freebsd-riscv64-v260306
---------- 2.7Mi -----------------------------------------------

.rwxr-xr-x 8.6Mi aba  6 Mar 16:27 gobrew-linux-arm64-v260306
.rwxr-xr-x 6.0Mi aba  6 Mar 17:57 gobrew-linux-arm64-v260306
---------- 2.6Mi -------------------------------------------

.rwxr-xr-x 9.4Mi aba  6 Mar 16:27 gobrew-linux-x64-v260306
.rwxr-xr-x 6.5Mi aba  6 Mar 17:57 gobrew-linux-x64-v260306
---------- 2.9Mi -----------------------------------------

.rwxr-xr-x 8.9Mi aba  6 Mar 16:29 gobrew-macos-apple-silicon-arm64-v260306
.rwxr-xr-x 6.1Mi aba  6 Mar 17:58 gobrew-macos-apple-silicon-arm64-v260306
---------- 2.8Mi ---------------------------------------------------------

.rwxr-xr-x 9.5Mi aba  6 Mar 16:29 gobrew-macos-intel-x64-v260306
.rwxr-xr-x 6.6Mi aba  6 Mar 17:58 gobrew-macos-intel-x64-v260306
---------- 2.9Mi -----------------------------------------------

.rwxr-xr-x 9.2Mi aba  6 Mar 16:30 gobrew-netbsd-amd64-v260306
.rwxr-xr-x 6.4Mi aba  6 Mar 17:59 gobrew-netbsd-amd64-v260306
---------- 2.8Mi --------------------------------------------

.rwxr-xr-x 8.5Mi aba  6 Mar 16:30 gobrew-netbsd-arm64-v260306
.rwxr-xr-x 5.9Mi aba  6 Mar 17:59 gobrew-netbsd-arm64-v260306
---------- 2.6Mi --------------------------------------------

.rwxr-xr-x 9.2Mi aba  6 Mar 16:30 gobrew-openbsd-amd64-v260306
.rwxr-xr-x 6.4Mi aba  6 Mar 17:59 gobrew-openbsd-amd64-v260306
---------- 2.8Mi ---------------------------------------------

.rwxr-xr-x 8.5Mi aba  6 Mar 16:30 gobrew-openbsd-arm64-v260306
.rwxr-xr-x 5.9Mi aba  6 Mar 18:00 gobrew-openbsd-arm64-v260306
---------- 2.6Mi ---------------------------------------------

.rwxr-xr-x 8.6Mi aba  6 Mar 16:28 gobrew-windows-arm64-v260306.exe
.rwxr-xr-x 6.0Mi aba  6 Mar 17:57 gobrew-windows-arm64-v260306.exe
---------- 2.6Mi -------------------------------------------------

.rwxr-xr-x 9.4Mi aba  6 Mar 16:28 gobrew-windows-x64-v260306.exe
.rwxr-xr-x 6.7Mi aba  6 Mar 17:57 gobrew-windows-x64-v260306.exe
---------- 2.7Mi -----------------------------------------------
```

It reduces the binary size by 2.7MiB in average which is +30% reduction in this case.
