package main

import (
	"flag"
	"log"

	"github.com/lomik/nooLiteHub/pkg/mtrf"
)

func main() {
	addr := flag.String("p", "/dev/ttyAMA0", "Serial port")
	// m := flag.String("m", "127.0.0.1:1883", "MQTT server")
	// t := flag.String("t", "nooLiteHub", "MQTT Root topic")

	device := mtrf.Connect(*addr)

	go func() {
		for {
			r := <-device.Recv()
			log.Println("mtrf received:\n", r)
		}
	}()

	select {}
}
