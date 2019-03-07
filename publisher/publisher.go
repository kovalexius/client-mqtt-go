package main

import (
	"os"
	"os/signal"
	"time"
	"flag"
	"fmt"
	"syscall"
	"crypto/tls"
	
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func publishing(client MQTT.Client) {
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		topic_context := "telemetry/"
		
		topic1 := topic_context + "light1"
		client.Publish(topic1, 0, false, t.String())
		
		topic2 := topic_context + "light2"
		client.Publish(topic2, 0, false, t.String())
		
		topic3 := topic_context + "light3"
		client.Publish(topic3, 0, false, t.String())
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	server := flag.String("server", "tcp://127.0.0.1:1883", "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")
	clientid := "publisher-test";
	
	connOpts := MQTT.NewClientOptions().AddBroker(*server).SetClientID(clientid).SetCleanSession(true)
	
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)
	
	connOpts.OnConnect = func(client MQTT.Client) {
		publishing(client)
	}
	
	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", *server)
	}
	
	<-c
}
