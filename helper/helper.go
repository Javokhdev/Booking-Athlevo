package helper

import (
	"database/sql"
	"time"
)

func DateToString(date sql.NullTime) string {
	if date.Valid {
		return date.Time.Format(time.RFC3339)
	}
	return ""
}
