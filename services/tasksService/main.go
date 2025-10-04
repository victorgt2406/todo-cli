package tasksService

import "gorm.io/gorm"

type TasksService struct {
	db *gorm.DB
}

func InitTaskService(db *gorm.DB) TasksService {
	return TasksService{
		db: db,
	}
}
