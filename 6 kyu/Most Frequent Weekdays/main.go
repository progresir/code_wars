package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	fmt.Println(MostFrequentDays(1984))
}

func MostFrequentDays(year int) []string {
	res := make([]string, 0)
	weekID := make([]int, 0)
	t := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	startWeek := t.Weekday()
	//println(time.Weekday(1).String())
	t = t.AddDate(1, 0, -1)
	for {
		if startWeek == t.Weekday() {
			weekID = append(weekID, int(startWeek))
			break
		}
		weekID = append(weekID, int(t.Weekday()))
		t = t.AddDate(0, 0, -1)
	}
	sort.Ints(weekID)
	sunday := false
	for i := range weekID {
		if weekID[i] == 0 {
			sunday = true
			continue
		}
		res = append(res, time.Weekday(weekID[i]).String())
	}
	if sunday {
		res = append(res, time.Weekday(0).String())
	}
	return res
}
