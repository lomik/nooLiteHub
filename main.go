package main

import (
	"flag"
	"log"
	"net"

	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"

	"github.com/lomik/nooLiteHub/pkg/mtrf"
)

func main() {
	addr := flag.String("p", "/dev/ttyAMA0", "Serial port")
	m := flag.String("m", "127.0.0.1:1883", "MQTT server")
	t := flag.String("t", "nooLiteHub", "MQTT Root topic")
	mqttClientID := flag.String("client", "nooLiteHub", "MQTT client ID")
	mqttUser := flag.String("user", "", "MQTT user")
	mqttPassword := flag.String("password", "", "MQTT password")

	device := mtrf.Connect(*addr)

	mqttConn, err := net.Dial("tcp", *m)
	if err != nil {
		log.Fatal(err)
		return
	}

	cc := mqtt.NewClientConn(conn)
	cc.Dump = false
	cc.ClientId = *mqttClientID

	tq := make([]proto.TopicQos, 1)
	tq[0].Topic = *t + "/#"
	// tq[0].Qos = proto.QosAtMostOnce

	if err := cc.Connect(*mqttUser, *mqttPassword); err != nil {
		log.Fatal(err)
		return
	}
	log.Println("connected with client id", cc.ClientId)
	cc.Subscribe(tq)

	go func() {
		for {
			r := <-device.Recv()
			log.Println("mtrf received:\n", r)
		}
	}()

	select {}
}
