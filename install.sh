#!/usr/bin/env bash
set -e

REPO="lgawin/now-cli"

if [ -n "$1" ]; then
  VERSION="$1"
  echo "🎯 Installing specified version: $VERSION"
else
  echo "🔍 Fetching the latest version tag..."
  VERSION=$(curl -sI "https://github.com/$REPO/releases/latest" | grep -Fi 'location:' | sed -E 's/.*\/tag\/([^[:space:]\r\n]+).*/\1/')

  if [ -z "$VERSION" ]; then
    echo "⚠️ Could not find a stable release via the API redirect."
    echo "💡 Defaulting to alpha bootstrapping version..."
    VERSION="v0.0.1-alpha"
  fi
fi
URL="https://github.com/$REPO/releases/download/$VERSION/now-linux-x86_64.tar.gz"
TARGET_DIR="/usr/local/bin"

echo "⏳ Starting now-cli installer ($VERSION)..."

if [ "$(id -u)" -ne 0 ]; then
  echo "🔐 Requesting sudo access to install to $TARGET_DIR..."
  curl -sL "$URL" | sudo tar -xz -C "$TARGET_DIR"
  sudo chmod +x "$TARGET_DIR"/now
  [ -f "$TARGET_DIR/adb-time-sync" ] && sudo chmod +x "$TARGET_DIR"/adb-time-sync
else
  curl -sL "$URL" | tar -xz -C "$TARGET_DIR"
  chmod +x "$TARGET_DIR"/now
  [ -f "$TARGET_DIR/adb-time-sync" ] && chmod +x "$TARGET_DIR"/adb-time-sync
fi

echo "✅ Installation complete!"
