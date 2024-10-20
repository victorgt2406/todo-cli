# Todo cli
A simple todo cli written in go

## Installation

```bash
go install github.com/victorgt2406/todo-cli@latest
```

## Usage

### Create a task
```bash
tdc go to the market
```

### View tasks WIP
List all relevant tasks
```bash
tdc l
```
List the important tasks
```bash
tdc li
```

### Delete tasks WIP
Remove all tasks
```bash
tdc ra
```
Remove task by id
```bash
tdc r[TASK ID]
```

Remove task by slug
```bash
tdc r [task-name] # ℹ️ help your self using tab ->
```

ℹ️ If you're using WSL2 and have Ollama native to Windows:
``` bash
# CMD
netsh interface portproxy add v4tov4 listenport=11434 listenaddress=0.0.0.0 connectport=11434 connectaddress=127.0.0.1
```
The only thing you need to do is use the private IP of your Windows machine.
