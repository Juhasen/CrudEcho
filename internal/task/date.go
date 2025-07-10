package task

import (
	"time"
)

const dateLayout = "02/01/2006"

func ParseDateStringToTime(dateStr string) (time.Time, error) {
	t, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		return time.Time{}, ErrInvalidDateFormat
	}
	return t, nil
}

func FormatTimeToDateString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(dateLayout)
}
