#!/usr/bin/env sh
# POSIX uninstaller for "gobrew" binary.
# Usage:
#   sh uninstall.sh
# or:
#   INSTALL_NAME="gobrew" INSTALL_DIR="/usr/local/bin" sh uninstall.sh
#
# If INSTALL_DIR is not given, it will:
#   1) Try `command -v i` to locate the binary.
#   2) Fallback to /usr/local/bin and /usr/bin.

set -e

INSTALL_NAME="${INSTALL_NAME:-gobrew}"

# Try to resolve current path of the binary.
FOUND_PATH="$(command -v "$INSTALL_NAME" 2>/dev/null || true)"

if [ -n "$INSTALL_DIR" ]; then
    # User explicitly specified install dir.
    TARGET="${INSTALL_DIR%/}/$INSTALL_NAME"
else
    if [ -n "$FOUND_PATH" ]; then
        TARGET="$FOUND_PATH"
    else
        # Fallback guesses: where the installer likely put it.
        if [ -x "/usr/local/bin/$INSTALL_NAME" ]; then
            TARGET="/usr/local/bin/$INSTALL_NAME"
        elif [ -x "/usr/bin/$INSTALL_NAME" ]; then
            TARGET="/usr/bin/$INSTALL_NAME"
        else
            echo "gobrew: not found in PATH and no INSTALL_DIR given." >&2
            exit 1
        fi
    fi
fi

# Normalize to remove trailing slashes.
case "$TARGET" in
    */) TARGET="${TARGET%/}/$INSTALL_NAME" ;;
esac

if [ ! -e "$TARGET" ]; then
    echo "Nothing to remove: $TARGET does not exist." >&2
    exit 0
fi

echo "This will remove:"
echo "  $TARGET"
printf "Proceed? [y/N] "
read ans || ans="n"

case "$ans" in
    y|Y|yes|YES)
        ;;
    *)
        echo "Aborted."
        exit 1
        ;;
esac

# Try to remove directly, then with sudo if needed.
if rm -f -- "$TARGET" 2>/dev/null; then
    :
else
    if command -v sudo >/dev/null 2>&1; then
        if sudo rm -f -- "$TARGET"; then
            :
        else
            echo "Failed to remove $TARGET (even with sudo)." >&2
            exit 1
        fi
    else
        echo "Cannot remove $TARGET (no permission and sudo not available)." >&2
        exit 1
    fi
fi

echo "Uninstalled $INSTALL_NAME from $TARGET"
