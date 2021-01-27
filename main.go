package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	_ = godotenv.Load()

	var listenSocket = flag.String("listen", ":2112", "Listen socket for exporter")
	var routerURL = flag.String("url", "", "DD-WRT router URL (make sure to add the final slash)")
	var username = flag.String("username", "", "Router user name")
	var password = flag.String("password", "", "Router password")
	var interfaces = flag.String("interfaces", "eth0,eth1,vlan0,br0", "Comma separated list of interface names for bandwidth scraping (eg. br0,vlan0,eth1)")

	flag.Parse()

	if *listenSocket == ":2112" && os.Getenv("DDWRT_LISTEN") != "" {
		*listenSocket = os.Getenv("DDWRT_LISTEN")
	}
	if *routerURL == "" {
		*routerURL = os.Getenv("DDWRT_URL")
	}
	if *username == "" {
		*username = os.Getenv("DDWRT_USERNAME")
	}
	if *password == "" {
		*password = os.Getenv("DDWRT_PASSWORD")
	}
	if *interfaces == "" {
		*interfaces = os.Getenv("DDWRT_INTERFACES")
	}

	if *routerURL == "" {
		fmt.Println("Please specify the router URL")
		os.Exit(1)
	}

	collector := NewWRTExporter(*routerURL, *username, *password, strings.Split(*interfaces, ","))

	registry := prometheus.NewRegistry()
	registry.MustRegister(collector)

	newHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	log.Println("Starting exporter at", *listenSocket)
	http.Handle("/metrics", promhttp.InstrumentMetricHandler(registry, newHandler))
	err := http.ListenAndServe(*listenSocket, nil)
	if err != nil {
		log.Fatal(err)
	}
}
