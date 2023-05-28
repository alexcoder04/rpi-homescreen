package data

import (
	"time"
)

func Time() map[string]any {
	now := time.Now()
	return map[string]any{
		"time": now.Format("15:04"),
		"date": now.Format("Monday 2"),
	}
}
