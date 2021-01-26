package main

import (
	"log"
	"strconv"
	"strings"
)

// MemInfo represent the memory info of the DD-WRT
type MemInfo struct {
	RAMTotal   int64
	RAMUsed    int64
	RAMFree    int64
	RAMShared  int64
	RAMBuffers int64
	RAMCached  int64
	SWAPTotal  int64
	SWAPUsed   int64
	SWAPFree   int64
}

// Scan is used to parse the formatted string from the DD-WRT router
func (wcl *MemInfo) Scan(value string) {
	var err error
	values := strings.Split(value, ",")
	if len(values) < 17 {
		return
	}
	wcl.RAMTotal, err = strconv.ParseInt(strings.Replace(values[8], "'", "", -1), 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RAMTotal %#v", err)
	}

	wcl.RAMUsed, err = strconv.ParseInt(strings.Replace(values[9], "'", "", -1), 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RAMUsed %#v", err)
	}

	wcl.RAMFree, err = strconv.ParseInt(strings.Replace(values[10], "'", "", -1), 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RAMFree %#v", err)
	}

	wcl.RAMShared, err = strconv.ParseInt(strings.Replace(values[11], "'", "", -1), 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RAMShared %#v", err)
	}

	wcl.RAMBuffers, err = strconv.ParseInt(strings.Replace(values[12], "'", "", -1), 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RAMBuffers %#v", err)
	}

	wcl.RAMCached, err = strconv.ParseInt(strings.Replace(values[13], "'", "", -1), 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: RAMCached %#v", err)
	}

	wcl.SWAPTotal, err = strconv.ParseInt(strings.Replace(values[15], "'", "", -1), 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: SWAPTotal %#v", err)
	}

	wcl.SWAPUsed, err = strconv.ParseInt(strings.Replace(values[16], "'", "", -1), 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: SWAPUsed %#v", err)
	}

	wcl.SWAPFree, err = strconv.ParseInt(strings.Replace(values[17], "'", "", -1), 10, 64)
	if err != nil {
		log.Printf("Error converting value to int: SWAPFree %#v", err)
	}
}
