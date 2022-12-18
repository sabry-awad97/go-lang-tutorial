package main

import (
	"fmt"
	"time"
)

// func PrintTime(label string, t *time.Time) {
// 	Printfln("%s: Day: %v: Month: %v Year: %v",
// 		label, t.Day(), t.Month(), t.Year())
// }

func representDateTime() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)
}

func PrintTime(label string, t *time.Time) {
	// layout := "Day: 02 Month: Jan Year: 2006"
	// fmt.Println(label, t.Format(layout))
	fmt.Println(label, t.Format(time.RFC822Z))
}

func FormattingTimeValues() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)
}

func parsingDateString() {
	layout := "2006-Jan-02"
	dates := []string{
		"1995-Jun-09",
		"2015-Jun-02",
	}
	for _, d := range dates {
		time, err := time.Parse(layout, d)
		if err == nil {
			PrintTime("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
}

func main() {
	// representDateTime()
	// FormattingTimeValues()
	parsingDateString()
}
