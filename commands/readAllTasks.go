package commands

import (
	"fmt"
	"todo-cli/models"
	t "todo-cli/services/tasksService"
	"todo-cli/utils"
)

func (command Command) readAllTasks(args []string) {
	tasks := command.tasksService.GetTasks(t.TaskFilter{}, t.OrderBy{})

	fmt.Println("Total tasks:", len(tasks))

	for _, task := range tasks {
		fmt.Println(strTask(task))
	}
}

func strTask(task models.Task) string {
	checked := " "
	strDate := ""
	if task.IsDone {
		checked = "x"
	}
	if task.TodoDate != nil {
		strDate = utils.FormatDateToString(*task.TodoDate)
	}
	return fmt.Sprintf("#%d - [%s] %s %s", task.ID, checked, task.Description, strDate)
}
