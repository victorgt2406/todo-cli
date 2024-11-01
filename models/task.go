package models

import "time"

type Task struct {
	ID          int
	Description string
	IsDone      bool
	Date        *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (t *Task) Default() {
	t.IsDone = false
	now := time.Now().UTC()
	t.CreatedAt = &now
	t.UpdatedAt = &now
}
