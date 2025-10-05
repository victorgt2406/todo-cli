# Todo cli
A simple todo cli written in go

## Installation
### Ubuntu
```bash
curl https://victorgt2406.github.io/todo-cli/install.sh | bash
```
### MacOS
```bash
curl https://victorgt2406.github.io/todo-cli/install.sh | zsh
```

#### Potential Installation Issue
The installation script tries to automatically append a line to your `~/.zshrc` file to add `$HOME/.local/bin` to your PATH, but it may fail due to insufficient write permissions on that file (common on macOS where .zshrc defaults to read-only like -rw-r--r--).

To resolve:
- Check permissions: `ls -l ~/.zshrc`
- Make the file writable if needed: `chmod u+w ~/.zshrc`
- Rerun the installation script.

Alternatively, manually add the following line to `~/.zshrc`:
```bash
export PATH="$HOME/.local/bin:$PATH"
```
Then reload your shell: `source ~/.zshrc`.

For more details, see:
- [Stack Overflow](https://stackoverflow.com/questions/64291625/zsh-permission-denied-users-macbookpro-zshrc)
- [Setapp](https://setapp.com/how-to/zsh-permissions-denied-error)
- [iBoysoft](https://iboysoft.com/howto/fix-zsh-permission-denied-in-mac-terminal.html)
- [Oh My Zsh Issue](https://github.com/ohmyzsh/ohmyzsh/issues/3291)

## Usage

### Create a task
```bash
tdc go to the market
```

### View tasks
```bash
tdc
```

## Configuration
### See the path runing the help command
``` bash
tdc help
```

### The application supports any LLM Provider that uses OpenAI API
```json
{
  "llm_provider": {
    "url": "",
    "model": "",
    "api_key": ""
  },
  "features": {
    "smart_task": false
  }
}
```
You will need to enable the smart task feature to use the LLM provider.