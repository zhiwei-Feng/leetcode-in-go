package n1185

import (
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	time := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return time.Weekday().String()
}