package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// WirelessClientList is the type for the list of the active wireless client
type WirelessClientList []WirelessClient

// WirelessClient is an active/connected wireless client
type WirelessClient struct {
	MAC       string
	Interface string
	Uptime    time.Duration
	TXRate    WiFiRate
	RXRate    WiFiRate
	Info      string
	Signal    int
	Noise     int
	SNR       int
	Quality   int // note that this value is between 0 and 1000
}

// Scan is used to parse the formatted string from the DD-WRT router
func (wcl *WirelessClientList) Scan(value string) {
	if strings.TrimSpace(value) == "" {
		return
	}

	value = strings.ReplaceAll(value, "day,", "day")
	value = strings.ReplaceAll(value, "days,", "day")
	values := strings.Split(value, ",")
	for i := 0; i*10 < len(values); i++ {
		var uptime time.Duration
		var err error
		var days int = 0
		var strduration string

		durationParts := strings.Split(strings.Trim(values[i*10+2], "'"), "day")
		if len(durationParts) == 1 {
			strduration = strings.Replace(strings.Replace(strings.Replace(durationParts[0], ":", "h", 1), ":", "m", 1), ":", "s", 1) + "s"
		} else {
			strduration = strings.Replace(strings.Replace(strings.Replace(durationParts[1], ":", "h", 1), ":", "m", 1), ":", "s", 1) + "s"
			days, _ = strconv.Atoi(strings.TrimSpace(durationParts[0]))
		}

		uptime, err = time.ParseDuration(strings.TrimSpace(strduration))
		if err != nil {
			log.Printf("Error converting value to uptime: %s %#v", strduration, err)
		}
		uptime += time.Duration(days*24) * time.Hour

		var txRate WiFiRate
		txRate.Scan(strings.Trim(values[i*10+3], "'"))

		var rxRate WiFiRate
		rxRate.Scan(strings.Trim(values[i*10+4], "'"))

		*wcl = append(*wcl, WirelessClient{
			MAC:       strings.Trim(values[i*10+0], "'"),
			Interface: strings.Trim(values[i*10+1], "'"),
			Uptime:    uptime,
			TXRate:    txRate,
			RXRate:    rxRate,
			Info:      strings.Trim(values[i*10+5], "'"),
			Signal:    MustAtoi(strings.Trim(values[i*10+6], "'")),
			Noise:     MustAtoi(strings.Trim(values[i*10+7], "'")),
			SNR:       MustAtoi(strings.Trim(values[i*10+8], "'")),
			Quality:   MustAtoi(strings.Trim(values[i*10+9], "'")),
		})
	}
}
