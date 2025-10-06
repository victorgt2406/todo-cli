package tasksService

import (
	"todo-cli/models"

	"gorm.io/gorm"
)

type TaskFilter struct {
	IsDone *bool
}

type OrderBy struct {
	DoneAsc      *bool
	CreatedAtAsc *bool
	TodoDateAsc  *bool
}

func (t TasksService) GetTasks(filter TaskFilter, orderBy OrderBy) []models.Task {
	var tasks []models.Task

	query := t.db.Model(&models.Task{})
	query = applyFilters(query, filter)
	query = applySorting(query, orderBy)
	query.Find(&tasks)

	return tasks
}

func applyFilters(query *gorm.DB, filter TaskFilter) *gorm.DB {
	if filter.IsDone != nil {
		query = query.Where("is_done = ?", *filter.IsDone)
	}
	return query
}

func applySorting(query *gorm.DB, orderBy OrderBy) *gorm.DB {
	if orderBy.DoneAsc != nil {
		query = query.Order("is_done " + ascDescToStr(*orderBy.DoneAsc))
	}
	if orderBy.TodoDateAsc != nil {
		query = query.Order("todo_date " + ascDescToStr(*orderBy.TodoDateAsc))
	}
	if orderBy.CreatedAtAsc != nil {
		query = query.Order("created_at " + ascDescToStr(*orderBy.CreatedAtAsc))
	}

	return query
}

func ascDescToStr(asc bool) string {
	if asc {
		return "ASC"
	} else {
		return "DESC"
	}
}
