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
### Database
The application supports two database providers:
- SQLite (default)
- PostgreSQL

You can configure the database provider in the `config.json` file:

```json
{
    "database": {
        "provider": "sqlite",  // or "postgres"
        "url": "data/todo.db" // SQLite file path or PostgreSQL connection URL
    }
}
```

For PostgreSQL, use the connection URL in the `path` field:
```json
{
    "database": {
        "provider": "postgres",
        "url": "postgres://username:password@localhost:5432/todo"
    }
}
```

The PostgreSQL URL format is: `postgres://username:password@host:port/dbname`
