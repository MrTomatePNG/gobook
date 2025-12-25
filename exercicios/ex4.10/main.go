package ex410

import (
	"fmt"
	"time"
)

// exercicio 4.10
func DateFormat(date time.Time) string {
	now := time.Now()
	diff := now.Sub(date)
	seconds := int(diff.Seconds())

	if seconds < 0 {
		return "Future"
	}
	switch {
	case seconds < 60:
		return "Now"
	case seconds < (60 * 60):
		minutes := seconds / 60
		return fmt.Sprintf("%d min ago", minutes)
	case seconds < (60 * 60 * 60):
		hours := seconds / (60 * 60)

		return fmt.Sprintf("%g hours", hours)
	case seconds < (60 * 60 * 60 * 30):
		days := seconds / (60 * 60 * 60)
		return fmt.Sprintf("%d days", days)
	case seconds < 31536000:
		months := seconds / 2592000
		if months <= 1 {
			return "More month"
		}
		return fmt.Sprintf("%d months", months)

	default:
		years := seconds / 31536000
		if years <= 1 {
			return "More year"
		}
		return fmt.Sprintf("%d years", years)
	}
}
