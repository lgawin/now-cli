#!/usr/bin/env bash
set -e

echo "🔍 Fetching the latest version of now-cli..."
VERSION=$(curl -s https://api.github.com/repos/lgawin/now-cli/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$VERSION" ]; then
    echo "❌ Error: Could not determine the latest release version." >&2
    exit 1
fi

URL="https://github.com/lgawin/now-cli/releases/download/$VERSION/now-linux-x86_64.tar.gz"
TARGET_DIR="/usr/local/bin"

echo "⏳ Starting now-cli installer ($VERSION)..."

if [ "$(id -u)" -ne 0 ]; then
    echo "🔐 Installation requires administrator privileges. Requesting sudo..."
    curl -sL "$URL" | sudo tar -xz -C "$TARGET_DIR"
    sudo chmod +x "$TARGET_DIR"/now
else
    curl -sL "$URL" | tar -xz -C "$TARGET_DIR"
    chmod +x "$TARGET_DIR"/now
fi

echo "✅ Installation complete!"
