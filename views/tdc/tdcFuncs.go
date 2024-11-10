package tdc

import (
	"todo-cli/models"

	"gorm.io/gorm"
)

func Check(m model) {
	m.tasks[m.cursor].IsDone = !m.tasks[m.cursor].IsDone
	m.db.Save(&m.tasks[m.cursor])
}

func findTasks(db *gorm.DB, tasks *[]models.Task) {
	db.Order("CASE WHEN date IS NULL THEN 0 ELSE 1 END").
		Order("date ASC").
		Order("id ASC").
		Find(&tasks)
}
