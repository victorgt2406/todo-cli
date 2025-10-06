package tasksPresenter

import "todo-cli/db"

func InitTasksPresenter(dbContext db.Context) TasksPresenter {
	return TasksPresenter{
		dbContext: dbContext,
	}
}
