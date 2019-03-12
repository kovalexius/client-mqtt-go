package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"./mongoDb"
	. "./conf"
)

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	mongoDb.InsertTelemetry(message.Topic(), string(message.Payload()))
}

func connect() {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered", r)
        }
    }()
	
	hostname, _ := os.Hostname()

	server := Conf.Mqtt.Server
	topic := Conf.Mqtt.Topic
	qos := Conf.Mqtt.Qos
	clientid := hostname + strconv.Itoa(time.Now().Second())
	if len(Conf.Mqtt.ClientId) > 0 {
		clientid = Conf.Mqtt.ClientId
	}
	username := Conf.Mqtt.Username
	password := Conf.Mqtt.Password

	connOpts := MQTT.NewClientOptions().AddBroker(server).SetClientID(clientid).SetCleanSession(true)
	if len(username) > 0 {
		connOpts.SetUsername(username)
		if len(password) > 0 {
			connOpts.SetPassword(password)
		}
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	connOpts.OnConnect = func(client MQTT.Client) {
		if token := client.Subscribe(topic, byte(qos), onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", server)
	}
}

func init() {
	connect()
}
