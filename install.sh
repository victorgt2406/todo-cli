#!/bin/bash

# Variables
TODO_CLI_VERSION="0.1.1"
RELEASE_URL="https://github.com/victorgt2406/todo-cli/archive/refs/tags/$TODO_CLI_VERSION.zip"
INSTALL_DIR="$HOME/.todo-cli"
BIN_DIR="$HOME/.local/bin"
CMD_NAME="tdc"

# Ensure ~/.local/bin is in PATH
if [[ ":$PATH:" != *":$BIN_DIR:"* ]]; then
    # Check if the PATH export line already exists in .bashrc
    if ! grep -q "export PATH=\$PATH:\$HOME/.local/bin" "$HOME/.bashrc"; then
        echo "Adding $BIN_DIR to PATH in ~/.bashrc"
        echo 'export PATH=$PATH:$HOME/.local/bin' >> "$HOME/.bashrc"
        source "$HOME/.bashrc"
    fi
fi

# Download and extract the release
echo "Downloading and installing todo-cli..."
mkdir -p "$INSTALL_DIR"
curl -L "$RELEASE_URL" -o "$INSTALL_DIR/release.zip"
if ! command -v unzip &> /dev/null; then
    echo "Error: 'unzip' command not found. Please install unzip first."
    exit 1
fi
unzip "$INSTALL_DIR/release.zip" -d "$INSTALL_DIR"
rm "$INSTALL_DIR/release.zip"

# Move contents from versioned folder to install dir
mv "$INSTALL_DIR/todo-cli-$TODO_CLI_VERSION"/* "$INSTALL_DIR/"
rm -r "$INSTALL_DIR/todo-cli-$TODO_CLI_VERSION"

# Compile the source code
echo "Compiling todo-cli..."
cd "$INSTALL_DIR"  # Change to the installation directory
if ! command -v go &> /dev/null; then
    echo "Error: 'go' command not found. Please install Go first."
    exit 1
fi
go build -o "$INSTALL_DIR/todo-cli" "$INSTALL_DIR/main.go"


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