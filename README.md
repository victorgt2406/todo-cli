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
### See the path
``` bash
tdc config # "tdc c" works too
```
### Database
The application supports two database providers:
- SQLite (default)
- PostgreSQL

You can configure the database provider in the `config.json` file:

```json
{
    "database": {
        "provider": "sqlite",  // or "postgres"
        "url": "todo-cli.db" // default sqlite path, but it can be where ever it is needed
    }
}
```

For PostgreSQL, use the connection URL in the `url` field:
```json
{
    "database": {
        "provider": "postgres",
        "url": "postgres://username:password@host:port/dbname"
    }
}
```
