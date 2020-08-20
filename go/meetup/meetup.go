package meetup

import (
	"time"
)

// WeekSchedule is an int representing a week schedule
type WeekSchedule int

// Define WeekSchedule
const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

// Day returns the nth day of the month of a specific year for a specific WeekSchedule and a particular Weekday
func Day(w WeekSchedule, wd time.Weekday, m time.Month, y int) int {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return 0
	}
	t := time.Date(y, m, 1, 0, 0, 0, 0, location)
	listWeekdays := []int{}
	for i := 0; i < 7; i++ {
		if t.AddDate(0, 0, i).Weekday() == wd {
			j := t.AddDate(0, 0, i)
			for j.Before(t.AddDate(0, 1, 0)) {
				listWeekdays = append(listWeekdays, j.Day())
				j = j.AddDate(0, 0, 7)
			}
		}
	}
	switch w {
	case First:
		return listWeekdays[0]
	case Second:
		return listWeekdays[1]
	case Third:
		return listWeekdays[2]
	case Fourth:
		return listWeekdays[3]
	case Last:
		return listWeekdays[len(listWeekdays)-1]
	case Teenth:
		if listWeekdays[1] >= 13 {
			return listWeekdays[1]
		}
		return listWeekdays[2]
	}
	return 0
}
