#!/bin/bash

echo "Welcome to the GPM installer"
echo "This script will download and install the GPM binary for your system."
echo "You might be prompted for your password to move the binary to /usr/local/bin."
echo ""
echo "Made by Nadhi.dev && pkg.lat"

# Detect OS and architecture
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

# Map uname output to expected values
case "$OS" in
    darwin)
        OS="darwin"
        ;;
    linux)
        OS="linux"
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

case "$ARCH" in
    x86_64|amd64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

URL="https://github.com/Nadhila-dot/GPM/raw/refs/heads/main/builds/gpm-$OS-$ARCH"

echo "Downloading GPM binary for $OS-$ARCH..."
curl -L "$URL" -o gpm || { echo "Download failed!"; exit 1; }

chmod +x gpm

echo "Installing GPM to /usr/local/bin (requires sudo)..."
sudo mv gpm /usr/local/bin/gpm || { echo "Install failed!"; exit 1; }

echo "GPM installed successfully! You can now run 'gpm' from anywhere."