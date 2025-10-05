package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string
	IsDone      bool
	TodoDate    *time.Time // pointer so it can be null
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) Default() {
	t.IsDone = false
}
