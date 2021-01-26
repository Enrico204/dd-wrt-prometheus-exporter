package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// InterfaceStatistics is the struct holding the interface counters
type InterfaceStatistics struct {
	RXBytes      int64
	RXPackets    int64
	RXErrors     int64
	RXDrop       int64
	RXFIFO       int64
	RXFrame      int64
	RXCompressed int64
	RXMulticast  int64

	TXBytes      int64
	TXPackets    int64
	TXErrors     int64
	TXDrop       int64
	TXFIFO       int64
	TXFrame      int64
	TXCompressed int64
	TXMulticast  int64
}

var cleanRx = regexp.MustCompile(`\s+`)

// Scan is used to parse the formatted string from the DD-WRT router
func (is *InterfaceStatistics) Scan(v string) {
	content := strings.SplitN(v, "\n", 2)
	if len(content) != 2 {
		log.Printf("Expected at least two lines, found %d instead", len(content))
		return
	}

	content[1] = cleanRx.ReplaceAllString(content[1], " ")
	rowvalues := strings.Split(strings.TrimSpace(content[1]), ":")
	if len(rowvalues) != 2 {
		log.Printf("Expected one semicolon, found %d instead", len(rowvalues)-1)
		return
	}

	var err error

	values := strings.Split(strings.TrimSpace(rowvalues[1]), " ")
	is.RXBytes, err = strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RXBytes %#v", err)
	}

	is.RXPackets, err = strconv.ParseInt(values[1], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RXPackets %#v", err)
	}

	is.RXErrors, err = strconv.ParseInt(values[2], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RXErrors %#v", err)
	}

	is.RXDrop, err = strconv.ParseInt(values[3], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RXDrop %#v", err)
	}

	is.RXFIFO, err = strconv.ParseInt(values[4], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RXFIFO %#v", err)
	}

	is.RXFrame, err = strconv.ParseInt(values[5], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RXFrame %#v", err)
	}

	is.RXCompressed, err = strconv.ParseInt(values[6], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RXCompressed %#v", err)
	}

	is.RXMulticast, err = strconv.ParseInt(values[7], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RXMulticast %#v", err)
	}

	is.TXBytes, err = strconv.ParseInt(values[8], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: TXBytes %#v", err)
	}

	is.TXPackets, err = strconv.ParseInt(values[9], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: TXPackets %#v", err)
	}

	is.TXErrors, err = strconv.ParseInt(values[10], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: TXErrors %#v", err)
	}

	is.TXDrop, err = strconv.ParseInt(values[11], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: TXDrop %#v", err)
	}

	is.TXFIFO, err = strconv.ParseInt(values[12], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: TXFIFO %#v", err)
	}

	is.TXFrame, err = strconv.ParseInt(values[13], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: TXFrame %#v", err)
	}

	is.TXCompressed, err = strconv.ParseInt(values[14], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: TXCompressed %#v", err)
	}

	is.TXMulticast, err = strconv.ParseInt(values[15], 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: TXMulticast %#v", err)
	}
}
