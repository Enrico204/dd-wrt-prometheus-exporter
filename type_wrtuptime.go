package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// WRTUptime is the struct that describe the uptime information from DD-WRT router
type WRTUptime struct {
	TimeNow     string
	UptimeSince time.Duration
	LoadAvg1m   float64
	LoadAvg5m   float64
	LoadAvg15m  float64
}

var uptimeRegex = regexp.MustCompile(`([0-9:]+) up ([0-9:]+),\s*load average: ([0-9\.]+), ([0-9\.]+), ([0-9\.]+)`)

// Scan is used to parse the formatted string from the DD-WRT router
func (up *WRTUptime) Scan(values string) {
	var err error
	matches := uptimeRegex.FindStringSubmatch(values)
	if len(matches) != 6 {
		return
	}

	up.TimeNow = matches[1]
	up.UptimeSince, err = time.ParseDuration(strings.Replace(matches[2], ":", "h", -1) + "m")
	if err != nil {
		log.Printf("Error converting value to uptime: %s %#v", matches[2], err)
	}

	up.LoadAvg1m, err = strconv.ParseFloat(matches[3], 64)
	if err != nil {
		log.Printf("Error converting value to float: %s %#v", matches[3], err)
	}

	up.LoadAvg5m, err = strconv.ParseFloat(matches[4], 64)
	if err != nil {
		log.Printf("Error converting value to float: %s %#v", matches[4], err)
	}

	up.LoadAvg15m, err = strconv.ParseFloat(matches[5], 64)
	if err != nil {
		log.Printf("Error converting value to float: %s %#v", matches[5], err)
	}
}
