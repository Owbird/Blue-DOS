package main

import (
	"log"
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	adapter.Enable()

	conns := []string{}

	adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		name := device.LocalName()

		if name != "" {
			go func() {
				log.Printf("Found device: %s", name)

				exists := false

				for _, addr := range conns {
					if addr == device.Address.String() {
						exists = true
						break
					}
				}

				if !exists {
					res, err := adapter.Connect(device.Address, bluetooth.ConnectionParams{
						ConnectionTimeout: bluetooth.NewDuration(time.Duration(time.Second * 1)),
						Timeout:           bluetooth.NewDuration(time.Duration(time.Second * 1)),
					})
					if err != nil {
						conns = append(conns, res.Address.String())
					}

				}
			}()
		}
	})
}
