package main

import (
	"reflect"
	"regexp"
)

var dataFormatRegex = regexp.MustCompile("{(.*?)::([^}]*)}")

// WirelessLive is the struct holding values from "Status_Wireless.live.asp" web API
type WirelessLive struct {
	MAC   string `wrt:"wl_mac"`
	ESSID string `wrt:"wl_ssid"`

	// TODO: 0 + 10 ???
	Channel string `wrt:"wl_channel"`

	// TODO: status to boolean
	RadioStatus string `wrt:"wl_radio"`

	// Radio TX in mW
	RadioTX *TXPowerType `wrt:"wl_xmit"`

	// Wi-Fi radio rate
	RadioRate *WiFiRate `wrt:"wl_rate"`

	ActiveWirelessClients *WirelessClientList `wrt:"active_wireless"`

	NetPacketInfo *PacketInfo `wrt:"packet_info"`

	Uptime *WRTUptime `wrt:"uptime"`
}

// Scan is used to parse the formatted string from the DD-WRT router
func (w *WirelessLive) Scan(content string) {
	t := reflect.TypeOf(*w)
	v := reflect.ValueOf(w).Elem()
	scanProtocol(content, t, v)
}
