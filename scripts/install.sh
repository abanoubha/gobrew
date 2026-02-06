#!/usr/bin/env sh
# POSIX installer for latest "i" release from GitHub.
# Supports Linux and macOS (Darwin) for x86_64/arm64.

set -e

GITHUB_REPO="abanoubha/gobrew"
INSTALL_NAME="gobrew"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

# --- detect OS/arch -> asset name ---

OS="$(uname -s 2>/dev/null || echo unknown)"
ARCH="$(uname -m 2>/dev/null || echo unknown)"

case "$OS" in
    Linux)
        case "$ARCH" in
            x86_64|amd64)   ASSET_NAME="gobrew-linux-x64" ;;
            aarch64|arm64)  ASSET_NAME="gobrew-linux-arm64" ;;
            *)
                echo "Unsupported Linux arch: $ARCH" >&2
                exit 1
                ;;
        esac
        ;;
    Darwin)
        case "$ARCH" in
            arm64)
                ASSET_NAME="gobrew-macos-apple-silicon-arm64"
                ;;
            x86_64|amd64)
                ASSET_NAME="gobrew-macos-intel-x64"
                ;;
            *)
                echo "Unsupported macOS arch: $ARCH" >&2
                exit 1
                ;;
        esac
        ;;
    Windows)
        case "$ARCH" in
            arm64)
                ASSET_NAME="gobrew-windows-arm64"
                ;;
            x86_64|amd64)
                ASSET_NAME="gobrew-windows-x64"
                ;;
            *)
                echo "Unsupported Windows arch: $ARCH" >&2
                exit 1
                ;;
        esac
        ;;
    *)
        echo "Unsupported OS: $OS" >&2
        exit 1
        ;;
esac

echo "Detected OS=$OS ARCH=$ARCH, selecting asset: $ASSET_NAME"

# --- HTTP helpers ---

if command -v curl >/dev/null 2>&1; then
    http_get() {
        curl -fsSL -H "Accept: application/vnd.github+json" \
                    -H "User-Agent: gobrew-latest-installer" \
                    "$1"
    }
    http_download() {
        curl -fsSL -H "User-Agent: gobrew-latest-installer" -o "$2" "$1"
    }
elif command -v wget >/dev/null 2>&1; then
    http_get() {
        wget -qO- --header="Accept: application/vnd.github+json" \
                  --header="User-Agent: gobrew-latest-installer" \
                  "$1"
    }
    http_download() {
        wget -qO "$2" --header="User-Agent: gobrew-latest-installer" "$1"
    }
else
    echo "ERROR: curl or wget required" >&2
    exit 1
fi

# --- query latest release and find matching asset ---

API_URL="https://api.github.com/repos/${GITHUB_REPO}/releases/latest"
RELEASE_JSON="$(http_get "$API_URL")" || {
    echo "ERROR: failed to get latest release for ${GITHUB_REPO}" >&2
    exit 1
}

# [debug] show JSON
# printf '%s\n' "$RELEASE_JSON" | awk '/"browser_download_url"/ { print $0 }' >&2

DOWNLOAD_URL=$(
    printf '%s\n' "$RELEASE_JSON" \
    | awk -v name="$ASSET_NAME" '
        {
            # find `"browser_download_url":"..."`
            while (match($0, /"browser_download_url"[[:space:]]*:[[:space:]]*"[^"]+"/)) {
                field = substr($0, RSTART, RLENGTH)
                # move past this match to look for the next one
                $0 = substr($0, RSTART + RLENGTH)

                # extract the URL inside the quotes
                gsub(/^.*"browser_download_url"[[:space:]]*:[[:space:]]*"/, "", field)
                gsub(/"$/, "", field)

                # match by substring on the asset name
                if (field ~ name) {
                    print field
                    exit 0
                }
            }
        }
    '
)

if [ -z "$DOWNLOAD_URL" ]; then
    echo "ERROR: asset \"$ASSET_NAME\" not found in latest release of ${GITHUB_REPO}" >&2
    exit 1
fi

echo "Downloading: $DOWNLOAD_URL"

# --- download and install ---

TMPDIR="${TMPDIR:-/tmp}"
TMPFILE="${TMPDIR}/.gobrew-latest-$$"

trap 'rm -f "$TMPFILE"' EXIT INT HUP TERM

http_download "$DOWNLOAD_URL" "$TMPFILE" || {
    echo "ERROR: download failed" >&2
    exit 1
}

# ensure install dir exists, maybe via sudo
if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creating ${INSTALL_DIR}..."
    if mkdir -p "$INSTALL_DIR" 2>/dev/null; then
        :
    elif command -v sudo >/dev/null 2>&1; then
        sudo mkdir -p "$INSTALL_DIR"
    else
        echo "ERROR: cannot create ${INSTALL_DIR}, try running with sudo" >&2
        exit 1
    fi
fi

TARGET="${INSTALL_DIR}/${INSTALL_NAME}"

if cp "$TMPFILE" "$TARGET" 2>/dev/null; then
    :
elif command -v sudo >/dev/null 2>&1; then
    sudo cp "$TMPFILE" "$TARGET"
else
    echo "ERROR: cannot write to ${INSTALL_DIR}, try running with sudo" >&2
    exit 1
fi

if chmod 755 "$TARGET" 2>/dev/null; then
    :
elif command -v sudo >/dev/null 2>&1; then
    sudo chmod 755 "$TARGET"
fi

echo "Installed ${INSTALL_NAME} to ${TARGET}"
echo "Done."
