package models

type ViewContext string

const (
	ViewNewTask  ViewContext = "newTask"
	ViewEditTask ViewContext = "editTask"
	ViewTasks    ViewContext = "tasks"
)
