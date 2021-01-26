package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	_ = godotenv.Load()

	var routerURL = flag.String("url", os.Getenv("DDWRT_URL"), "DD-WRT router URL (make sure to add the final slash)")
	var username = flag.String("username", os.Getenv("DDWRT_USERNAME"), "Router user name")
	var password = flag.String("password", os.Getenv("DDWRT_PASSWORD"), "Router password")
	var interfaces = flag.String("interfaces", os.Getenv("DDWRT_INTERFACES"), "Comma separated list of interface names for bandwidth scraping (eg. br0,vlan0,eth1)")

	if *routerURL == "" {
		fmt.Println("Please specify the router URL")
		os.Exit(1)
	}

	collector := NewWRTExporter(*routerURL, *username, *password, strings.Split(*interfaces, ","))

	registry := prometheus.NewRegistry()
	registry.MustRegister(collector)

	newHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	http.Handle("/metrics", promhttp.InstrumentMetricHandler(registry, newHandler))
	http.ListenAndServe(":2112", nil)
}
