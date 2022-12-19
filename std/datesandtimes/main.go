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

func parseLocation() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nycerr := time.LoadLocation("America/New_York")
	local, _ := time.LoadLocation("Local")

	if lonerr == nil && nycerr == nil {
		nolocation, _ := time.Parse(layout, date)
		londonTime, _ := time.ParseInLocation(layout, date, london)
		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
		localTime, _ := time.ParseInLocation(layout, date, local)
		PrintTime("No location:", &nolocation)
		PrintTime("London:", &londonTime)
		PrintTime("New York:", &newyorkTime)
		PrintTime("Local:", &localTime)
	} else {
		fmt.Println(lonerr.Error(), nycerr.Error())
	}
}

func specifyTimeZone() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london := time.FixedZone("BST", 1*60*60)
	newyork := time.FixedZone("EDT", -4*60*60)
	local := time.FixedZone("Local", 0)

	nolocation, _ := time.Parse(layout, date)
	londonTime, _ := time.ParseInLocation(layout, date, london)
	newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
	localTime, _ := time.ParseInLocation(layout, date, local)
	PrintTime("No location:", &nolocation)
	PrintTime("London:", &londonTime)
	PrintTime("New York:", &newyorkTime)
	PrintTime("Local:", &localTime)
}

func main() {
	// representDateTime()
	// FormattingTimeValues()
	// parsingDateString()
	// parseLocation()
	specifyTimeZone()
}
