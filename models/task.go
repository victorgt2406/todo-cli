package models

import "time"

type Task struct {
	ID          int
	Description string
	IsDone      bool
	Date        string
	CreatedAt   string
	UpdatedAt   string
}

func (t *Task) Default() {
	t.IsDone = false
	t.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	t.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
}
