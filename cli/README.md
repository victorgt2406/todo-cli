# Todo CLI Architecture

This CLI application follows a Model-View-Controller (MVC) pattern using the Bubble Tea framework:

## Architecture Components

### Model (`model.go`)
The **Model** represents the state and data needed to make the CLI work. It contains:
- Database connection (`db`)
- Task data (`tasks`)
- UI state (`cursor`, `viewContext`, `textInput`)
- Database context (`dbContext`)

The model is the single source of truth for the application's state.

### View (`view.go`)
The **View** handles the presentation layer - what the user sees on screen. It includes:
- `View()` - Main view function that renders the current UI
- `viewTasks()` - Renders the task list with cursor navigation
- Handles different view contexts (task list, new task, edit task)

### Controller (`update.go`)
The **Update** function acts as the controller, handling user input and business logic:
- Processes keyboard events and user interactions
- Manages state transitions between different views
- Delegates to specific update handlers (`updateTask`, `updateTasks`)
- Handles application lifecycle (quit on Ctrl+C)

### Initialization (`init.go`)
The **Init** component handles starting the CLI with default values:
- `Start()` - Bootstraps the application with database and context
- `initialModel()` - Creates the initial state with default values
- Loads existing tasks from the database
- Sets up the initial view context and text input

## Data Flow

1. **Init** → Sets up initial state and starts the application
2. **Model** → Maintains current state and data
3. **Update** → Processes user input and updates the model
4. **View** → Renders the updated state to the user
5. Repeat steps 3-4 for user interactions
