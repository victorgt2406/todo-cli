package controllers

import (
	"todo-cli/configs"
	"todo-cli/models"

	"gorm.io/gorm"
)

type TaskController struct {
	db *gorm.DB
}

// TASK Controller CRUD

func NewTaskController() *TaskController {
	db := configs.InitDB()
	return &TaskController{db: db}
}

func NewTaskControllerWithDB(db *gorm.DB) *TaskController {
	return &TaskController{db: db}
}

func (c *TaskController) GetAllTasks() []models.Task {
	var tasks []models.Task
	c.db.Find(&tasks)
	return tasks
}

func (c *TaskController) GetTask(id int) models.Task {
	var task models.Task
	c.db.First(&task, id)
	return task
}

func (c *TaskController) CreateTask(task models.Task) int {
	c.db.Create(&task)
	return task.ID
}

func (c *TaskController) UpdateTask(id int, task models.Task) {
	c.db.Model(&task).Where("id = ?", id).Updates(task)
}

func (c *TaskController) DeleteTask(id int) {
	c.db.Delete(&models.Task{}, id)
}
