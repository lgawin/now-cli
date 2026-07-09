#!/usr/bin/env bash
set -e

VERSION="v0.0.1-alpha"
URL="https://github.com/lgawin/now-cli/releases/download/$VERSION/now-linux-x86_64.tar.gz"
TARGET_DIR="/usr/local/bin"

if [ "$(id -u)" -ne 0 ]; then
    echo "🔐 Installation requires administrator privileges. Requesting sudo..."
    curl -sL "$URL" | sudo tar -xz -C "$TARGET_DIR"
    sudo chmod +x "$TARGET_DIR"/now
else
    curl -sL "$URL" | tar -xz -C "$TARGET_DIR"
    chmod +x "$TARGET_DIR"/now
fi

echo "✅ Installation complete!"
