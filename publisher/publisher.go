package main

import (
	"os"
	"os/signal"
	"time"
	"fmt"
	"math/rand"
	"syscall"
	"crypto/tls"
	"strconv"
	
	MQTT "github.com/eclipse/paho.mqtt.golang"
	. "./conf"
)

func publishing(client MQTT.Client) {
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		topic_context := Conf.TopicContext
		
		i := rand.Intn(len(Conf.Topics))
		topic := topic_context + Conf.Topics[i]
		client.Publish(topic, 0, false, strconv.FormatInt(t.Unix(), 10))
		rand.Seed(time.Now().Unix())
	}
}

func main() {
	//MQTT.DEBUG = log.New(os.Stdout, "", 0)
	//MQTT.ERROR = log.New(os.Stdout, "", 0)
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	server := Conf.Server
	clientid := Conf.ClientId;
	
	connOpts := MQTT.NewClientOptions().AddBroker(server).SetClientID(clientid).SetCleanSession(true)
	
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)
	
	connOpts.OnConnect = func(client MQTT.Client) {
		publishing(client)
	}
	
	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", server)
	}
	
	<-c
}
