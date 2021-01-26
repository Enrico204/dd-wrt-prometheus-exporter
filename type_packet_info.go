package main

import (
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var packetInfoRegex = regexp.MustCompile("(.*?)=([0-9]*)")

// PacketInfo is the struct holding values for TX/RX packets stats
type PacketInfo struct {
	RXGoodPacket  int64 `wrturl:"SWRXgoodPacket"`
	RXErrorPacket int64 `wrturl:"SWRXerrorPacket"`
	TXGoodPacket  int64 `wrturl:"SWTXgoodPacket"`
	TXErrorPacket int64 `wrturl:"SWTXerrorPacket"`
}

// Scan is used to parse the formatted string from the DD-WRT router
func (pi *PacketInfo) Scan(values string) {
	t := reflect.TypeOf(*pi)
	v := reflect.ValueOf(pi).Elem()

	matches := packetInfoRegex.FindAllStringSubmatch(values, -1)
	for _, m := range matches {
		for i := 0; i < t.NumField(); i++ {
			tag := t.Field(i).Tag.Get("wrturl")
			f := v.Field(i)
			if tag == strings.TrimPrefix(m[1], ";") && f.CanSet() {
				intval, err := strconv.ParseInt(m[2], 10, 64)
				if err != nil {
					log.Printf("Error converting value to int: %s %#v", tag, err)
				} else {
					f.SetInt(intval)
				}
			}
		}
	}
}
