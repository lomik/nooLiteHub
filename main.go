package main

import (
	"flag"
	"log"

	"github.com/lomik/nooLiteHub/pkg/hub"
	"github.com/lomik/nooLiteHub/pkg/mtrf"
)

func onMessage(topic string, payload []byte) {
	log.Println("on message", topic, string(payload))
}

func main() {
	port := flag.String("port", "/dev/ttyAMA0", "Serial port")
	broker := flag.String("server", "127.0.0.1:1883", "MQTT broker")
	topic := flag.String("topic", "nooLiteHub", "MQTT root topic")
	mqttClientID := flag.String("client", "nooLiteHub", "MQTT client ID")
	mqttUser := flag.String("user", "", "MQTT user")
	mqttPassword := flag.String("password", "", "MQTT password")

	device := mtrf.Connect(*port)

	h, err := hub.New(device, hub.Options{
		Broker:   *broker,
		Topic:    *topic,
		ClientID: *mqttClientID,
		User:     *mqttUser,
		Password: *mqttPassword,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	h.Loop()
}
