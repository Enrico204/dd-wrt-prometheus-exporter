package main

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

// CustomScanner is the interface for DD-WRT scannable structs
type CustomScanner interface {
	Scan(string)
}

// MustAtoi forces the number conversion to integer using Atoi
func MustAtoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int(v)
}

// TXPowerType is the type for Wi-Fi TX power
type TXPowerType int

// Scan is used to parse the formatted string from the DD-WRT router
func (t *TXPowerType) Scan(v string) {
	if v == "Radio is Off" {
		*t = 0
		return
	}

	intval, err := strconv.Atoi(strings.ReplaceAll(v, " mW", ""))
	if err != nil {
		log.Printf("Error converting TX power: %#v", err)
	} else {
		*t = TXPowerType(intval)
	}
}

// WiFiRate is the type for Wi-Fi rate
// TODO add constants?
type WiFiRate float32

// Scan is used to parse the formatted string from the DD-WRT router
func (t *WiFiRate) Scan(v string) {
	intval, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(v, " Mbps", ""), "M", ""), 32)
	if err != nil {
		log.Printf("Error converting Wi-Fi rate: %#v", err)
	} else {
		*t = WiFiRate(intval)
	}
}

func scanProtocol(content string, t reflect.Type, v reflect.Value) {

	customScannerType := reflect.TypeOf((*CustomScanner)(nil)).Elem()

	matches := dataFormatRegex.FindAllStringSubmatch(string(content), -1)
	for _, m := range matches {
		for i := 0; i < t.NumField(); i++ {
			tag := t.Field(i).Tag.Get("wrt")

			f := v.Field(i)
			if tag == m[1] && f.CanSet() {
				if f.Type().Implements(customScannerType) {
					val := reflect.New(f.Type().Elem())
					inputs := []reflect.Value{
						reflect.ValueOf(m[2]),
					}
					val.MethodByName("Scan").Call(inputs)
					f.Set(val)
				} else if f.Kind() == reflect.Int {
					intval, err := strconv.ParseInt(m[2], 10, 64)
					if err != nil {
						log.Printf("Error converting value to int: %s %#v", tag, err)
					} else {
						f.SetInt(intval)
					}
				} else if f.Kind() == reflect.Float32 {
					fval, err := strconv.ParseFloat(m[2], 32)
					if err != nil {
						log.Printf("Error converting value to float32: %s %#v", tag, err)
					} else {
						f.SetFloat(fval)
					}
				} else if f.Kind() == reflect.Float64 {
					fval, err := strconv.ParseFloat(m[2], 64)
					if err != nil {
						log.Printf("Error converting value to float64: %s %#v", tag, err)
					} else {
						f.SetFloat(fval)
					}
				} else if f.Kind() == reflect.String {
					f.SetString(m[2])
					/*} else {
					fmt.Println(f.Type().Name())
					fmt.Println(f.Type().Implements(customScannerType))*/
				}
			}
		}
	}
}
