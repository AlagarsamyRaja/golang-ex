package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now()
	fmt.Println(date.Format("02-01-2006"))

	numericMonth := int(date.Month())
	fmt.Println("Month in numeric:", numericMonth)

	stringMonth := date.Month()
	fmt.Println("Month in word:", stringMonth)

	targetWeekDay := time.Wednesday
	for i := 7; i > 0; i-- {
		days := date.AddDate(0, 0, -i)
		if days.Weekday() == targetWeekDay {
			fmt.Println("Last Wednesday in This month:", days.Format("02-01-2006"))
			break
		}
	}

	firstDay := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	lastDay := firstDay.AddDate(0, 1, -1)
	fmt.Println("Last Day in this month:", lastDay.Format("Monday"))

	fmt.Println("Last 5 days in this month")
	for i := -1; i >= -5; i-- {
		dates := date.AddDate(0, 0, i)
		fmt.Println(dates.Format("02-Jan-2006 - Monday"))
	}

	now := time.Now()
	var seconds int
	var defaultTime time.Time
	fmt.Print("Input Seconds:")
	fmt.Scanln(&seconds)

	if seconds < 0 {
		seconds = seconds * -1
	}
	days := seconds / 86400
	remainingSeconds := seconds % 86400
	hours := remainingSeconds / 3600
	remainingSeconds = seconds % 3600
	minutes := remainingSeconds / 60
	seconds = remainingSeconds % 60

	defaultTime = defaultTime.AddDate(0, 0, days).Add(time.Duration(hours) * time.Hour).Add(time.Duration(minutes) * time.Minute).Add(time.Duration(seconds) * time.Second)
	fmt.Println("Adding hours,minutes,seconds in default time:", defaultTime)

	now = now.AddDate(0, 0, days).Add(time.Duration(hours) * time.Hour).Add(time.Duration(seconds) * time.Second).Add(time.Duration(minutes) * time.Minute)
	fmt.Println("Adding hours,minutes,seconds in current time:", now)

	fmt.Printf("D:H:M:S - %d:%d:%d:%d", days, hours, minutes, seconds)
}
