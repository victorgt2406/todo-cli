package tasksPresenter

import "todo-cli/db"

type TasksPresenter struct {
	dbContext db.Context
}

type viewContext string

const (
	viewNewTask  viewContext = "newTask"
	viewEditTask viewContext = "editTask"
	viewTasks    viewContext = "tasks"
)
