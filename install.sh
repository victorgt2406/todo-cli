#!/bin/bash

# Variables
TODO_CLI_VERSION="0.1"
RELEASE_URL="https://github.com/victorgt2406/todo-cli/archive/refs/tags/$TODO_CLI_VERSION.zip"
INSTALL_DIR="$HOME/.todo-cli"
BIN_DIR="$HOME/.local/bin"
CMD_NAME="tdc"

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

# Unset environment variables used during installation
unset INSTALL_DIR
unset BIN_DIR
unset CMD_NAME
unset RELEASE_URL
unset TODO_CLI_VERSION
unset SHELL_CONFIG
