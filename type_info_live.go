package main

import (
	"reflect"
)

// InfoLive is the struct holding values from "Info.live.htm" web API
type InfoLive struct {
	LanMAC string `wrt:"lan_mac"`
	WanMAC string `wrt:"wan_mac"`
	MAC    string `wrt:"wl_mac"`
	LanIP  string `wrt:"lan_ip"`

	// TODO: 0 + 10 ???
	Channel string `wrt:"wl_channel"`

	// TODO: status to boolean
	RadioStatus string `wrt:"wl_radio"`

	// Radio TX in mW
	RadioTX *TXPowerType `wrt:"wl_xmit"`

	// Wi-Fi radio rate
	RadioRate *WiFiRate `wrt:"wl_rate"`

	NetPacketInfo *PacketInfo `wrt:"packet_info"`

	WirelessModeShort string `wrt:"wl_mode_short"`

	LanProto string `wrt:"lan_proto"`

	MemInfo *MemInfo `wrt:"mem_info"`

	ActiveWirelessClients *WirelessClientList `wrt:"active_wireless"`

	Uptime *WRTUptime `wrt:"uptime"`
}

// Scan is used to parse the formatted string from the DD-WRT router
func (w *InfoLive) Scan(content string) {
	t := reflect.TypeOf(*w)
	v := reflect.ValueOf(w).Elem()
	scanProtocol(content, t, v)
}
