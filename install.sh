#!/bin/bash

# Variables
TODO_CLI_VERSION="0.1"
INSTALL_DIR="$HOME/.todo-cli"
BIN_DIR="$HOME/.local/bin"
CMD_NAME="tdc"

# Detect system architecture and OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture names
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

# Map OS names
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
RELEASE_URL="https://github.com/victorgt2406/todo-cli/releases/download/v${TODO_CLI_VERSION}/${BINARY_NAME}"

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
    fi
fi

# Create installation directory
mkdir -p "$INSTALL_DIR"

# Download the binary
echo "Downloading todo-cli binary..."
curl -L "$RELEASE_URL" -o "$INSTALL_DIR/todo-cli"

# Ensure the binary is executable
chmod +x "$INSTALL_DIR/todo-cli"

# Link to a binary in the user's local bin
mkdir -p "$BIN_DIR"
ln -sf "$INSTALL_DIR/todo-cli" "$BIN_DIR/$CMD_NAME"

# Ensure the new PATH is available in the current session
export PATH="$PATH:$BIN_DIR"

echo "Installation complete. You can now use the command '$CMD_NAME'."

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
unset RELEASE_URL
unset TODO_CLI_VERSION
unset SHELL_CONFIG
unset OS
unset ARCH
unset BINARY_NAME
