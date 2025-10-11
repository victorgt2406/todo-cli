package commands

import (
	"fmt"
)

func (c Command) checkTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a task ID to check.")
		return
	}
	taskID := args[0]
	task, err := c.tasksService.GetTaskByID(taskID)
	if err != nil {
		fmt.Printf("Task not found. Please make sure to enter a registered task ID.\n")
		return
	}

	task.IsDone = !task.IsDone
	c.tasksService.UpdateTask(task)

	if task.IsDone {
		fmt.Printf("✓ Task %s is now done.\n", taskID)
	} else {
		fmt.Printf("○ Task %s is now pending.\n", taskID)
	}
}
