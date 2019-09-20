package main

import (
	"bytes"
	"flag"
	"log"
	"net"

	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"

	"github.com/lomik/nooLiteHub/pkg/mtrf"
)

func main() {
	port := flag.String("port", "/dev/ttyAMA0", "Serial port")
	server := flag.String("server", "127.0.0.1:1883", "MQTT server")
	topic := flag.String("topic", "nooLiteHub", "MQTT Root topic")
	mqttClientID := flag.String("client", "nooLiteHub", "MQTT client ID")
	mqttUser := flag.String("user", "", "MQTT user")
	mqttPassword := flag.String("password", "", "MQTT password")

	device := mtrf.Connect(*port)

	mqttConn, err := net.Dial("tcp", *server)
	if err != nil {
		log.Fatal(err)
		return
	}

	cc := mqtt.NewClientConn(mqttConn)
	cc.Dump = false
	cc.ClientId = *mqttClientID

	tq := make([]proto.TopicQos, 1)
	tq[0].Topic = *topic + "/write/#"
	tq[0].Qos = proto.QosAtMostOnce

	if err := cc.Connect(*mqttUser, *mqttPassword); err != nil {
		log.Fatal(err)
		return
	}
	log.Println("connected with client id", cc.ClientId)
	cc.Subscribe(tq)

	mqttSend := func(t string, m string) {
		log.Printf("[mqtt send] %s: %s", *topic+t, m)
		cc.Publish(&proto.Publish{
			Header:    proto.Header{},
			TopicName: *topic + t,
			Payload:   proto.BytesPayload([]byte(m)),
		})
	}

	go func() {
		for {
			r := <-device.Recv()
			mqttSend("/in/raw", r.String())
		}
	}()

	for m := range cc.Incoming {
		b := new(bytes.Buffer)
		m.Payload.WritePayload(b)
		log.Printf("[mqtt recv] %s: %s", m.TopicName, b.String())
		// onMessage([]byte(m.TopicName), b.Bytes())
	}
}
