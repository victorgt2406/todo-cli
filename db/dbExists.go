package db

import "os"

func dbExists(dbPath string) bool {
	if _, err := os.Stat(dbPath); err == nil {
		return true
	}
	return false
}
