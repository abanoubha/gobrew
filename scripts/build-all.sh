#!/usr/bin/env sh
# build-all.sh – cross-compile Go binaries for multiple platforms

# Fail fast on errors
set -eu

# Project to build
PACKAGE_PATH="."

# Base name of the binary (without OS/arch/extension/version)
BASE_NAME="gobrew"

# List of target platforms as "GOOS/GOARCH"
# go tool dist list
PLATFORMS="
linux/amd64
linux/arm64
windows/amd64
windows/arm64
darwin/arm64
darwin/amd64
freebsd/amd64
freebsd/arm64
freebsd/riscv64
netbsd/amd64
netbsd/arm64
openbsd/amd64
openbsd/arm64
"

OUT_DIR="./dist"

mkdir -p "$OUT_DIR"

echo "Building Go binaries into: $OUT_DIR"

for PLATFORM in $PLATFORMS; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}

    OUT_OS="$GOOS"
    OUT_ARCH="$GOARCH"

    case "$GOOS" in
        linux)
            OUT_OS="linux"
                case "$GOARCH" in
                    amd64)
                        OUT_ARCH="x64"
                        ;;
                    arm64)
                        OUT_ARCH="arm64"
                        ;;
                esac
            ;;
        windows)
            OUT_OS="windows"
                case "$GOARCH" in
                    amd64)
                        OUT_ARCH="x64"
                        ;;
                    arm64)
                        OUT_ARCH="arm64"
                        ;;
                esac
            ;;
        darwin)
            OUT_OS="macos"
                case "$GOARCH" in
                amd64)
                    OUT_ARCH="intel-x64"
                    ;;
                arm64)
                    OUT_ARCH="apple-silicon-arm64"
                    ;;
            esac
            ;;
    esac

    OUTPUT_NAME="$BASE_NAME-$OUT_OS-$OUT_ARCH"

    if [ -n "$1" ]; then
        OUTPUT_NAME="$OUTPUT_NAME-$1"
    fi

    if [ "$GOOS" = "windows" ]; then
        OUTPUT_NAME="$OUTPUT_NAME.exe"
    fi

    echo "==> Building for GOOS=$GOOS GOARCH=$GOARCH -> $OUTPUT_NAME"

    GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$OUT_DIR/$OUTPUT_NAME" "$PACKAGE_PATH"
done

echo "Done."
