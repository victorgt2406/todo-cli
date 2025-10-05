#!/bin/bash

# todo-cli install script

set -e
set -u
set -o pipefail

# Variables
TODO_CLI_VERSION=$(curl -s https://api.github.com/repos/victorgt2406/todo-cli/releases/latest | grep -o '"tag_name": "[^"]*' | grep -o '[^"]*$')
INSTALL_DIR="$HOME/.todo-cli"
CMD_NAME="tdc"
BIN_DIR="$HOME/.local/bin"
CONFIG_DIR="$INSTALL_DIR"
TEMP_DIR="/tmp/todo-cli-install"
RELEASE_URL_SOURCE="https://github.com/victorgt2406/todo-cli/archive/refs/tags/${TODO_CLI_VERSION}.zip"

# Detect system architecture and OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

case $OS in
    linux)
        OS="linux"
        ;;
    darwin)
        OS="darwin"
        ;;
    *)
        echo "Unsupported operating system: $OS"
        exit 1
        ;;
esac

BINARY_NAME="todo-cli-${OS}-${ARCH}"
RELEASE_URL_BIN="https://github.com/victorgt2406/todo-cli/releases/download/${TODO_CLI_VERSION}/${BINARY_NAME}"

# Ensure ~/.local/bin is in PATH
if [[ ":$PATH:" != *":$BIN_DIR:"* ]]; then
    # Determine shell configuration file
    SHELL_CONFIG=""
    if [ -n "$ZSH_VERSION" ]; then
        SHELL_CONFIG="$HOME/.zshrc"
    elif [ -n "$BASH_VERSION" ]; then
        SHELL_CONFIG="$HOME/.bashrc"
    fi

    if [ -n "$SHELL_CONFIG" ]; then
        # Check if the PATH export line already exists
        if ! grep -q "export PATH=\$PATH:\$BIN_DIR" "$SHELL_CONFIG"; then
            echo "Adding $BIN_DIR to PATH in $SHELL_CONFIG"
            echo 'export PATH=$PATH:$BIN_DIR' >> "$SHELL_CONFIG"
            source "$SHELL_CONFIG"
        fi
    else
        echo "Warning: Unable to determine shell configuration file"
        echo "Please manually add the following line to your shell configuration file:"
        echo "  export PATH=\$PATH:$BIN_DIR"
    fi
fi
echo "Downloading todo-cli version (v$TODO_CLI_VERSION)"
echo "release URL: $RELEASE_URL_BIN"
echo "Using architecture: $ARCH with the OS: $OS and the binary name: $BINARY_NAME"

# Clean up previous installation (with safety checks)
if [ -n "$BIN_DIR" ] && [ -n "$CMD_NAME" ]; then
    rm -f "$BIN_DIR/$CMD_NAME"
fi

if [ -n "$INSTALL_DIR" ] && [ "$INSTALL_DIR" != "$HOME" ] && [ "$INSTALL_DIR" != "/" ]; then
    rm -rf "$INSTALL_DIR"
fi

# Create necessary directories
mkdir -p "$INSTALL_DIR"
mkdir -p "$CONFIG_DIR"
mkdir -p "$TEMP_DIR"

# Download binary
echo "Downloading todo-cli binary..."
curl -L "$RELEASE_URL_BIN" -o "$INSTALL_DIR/todo-cli"

# Ensure the binary is executable
chmod +x "$INSTALL_DIR/todo-cli"

# Link to a binary in the user's local bin
mkdir -p "$BIN_DIR"
ln -sf "$INSTALL_DIR/todo-cli" "$BIN_DIR/$CMD_NAME"

# Ensure the new PATH is available in the current session
export PATH="$PATH:$BIN_DIR"

echo "Installation completed."

# Check for successful installation
if command -v "$CMD_NAME" &> /dev/null; then
    echo "$CMD_NAME is now available on your PATH!"
else
    echo "Something went wrong. Please check your PATH settings."
fi

# Unset environment variables used during installation
unset INSTALL_DIR
unset BIN_DIR
unset CMD_NAME
unset RELEASE_URL_BIN
unset RELEASE_URL_SOURCE
unset TODO_CLI_VERSION
unset SHELL_CONFIG
unset OS
unset ARCH
unset BINARY_NAME

# Make the command available in current session
if [ -n "$ZSH_VERSION" ]; then
    source ~/.zshrc
elif [ -n "$BASH_VERSION" ]; then
    source ~/.bashrc
else
    echo "Please restart your terminal or run 'source ~/.bashrc' (or equivalent) to use the command."
fi
