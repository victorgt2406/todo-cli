package utils

import (
	"fmt"
	"time"
)

func FormatDateToString(t time.Time) string {
	weekday := t.Weekday().String()
	day := t.Day()
	month := t.Month().String()
	return fmt.Sprintf("%s %d of %s", weekday, day, month)
}
