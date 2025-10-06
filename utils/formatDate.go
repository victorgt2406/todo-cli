package utils

import (
	"fmt"
	"time"
)

func FormatDateToString(t *time.Time) *string {
	if t == nil {
		return nil
	}
	weekday := t.Weekday().String()
	day := t.Day()
	month := t.Month().String()
	result := fmt.Sprintf("%s %d of %s", weekday, day, month)
	return &result
}
