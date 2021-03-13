package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"keplerhub/pkg/megad2561"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	options := megad2561.MegadServiceOptions{
		HostIP: "192.168.88.14",
		Pwd:    "sec",
	}
	service := megad2561.NewMegadService(options)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Initialized config:")
	fmt.Println(service.GetConfig())

	err := service.BatchConfigure(func(builder megad2561.BuilderConfig) megad2561.BuilderConfig {
		if !service.HasID() {
			err := builder.SetID("megad")
			if err != nil {
				fmt.Println(err)
			}
		}

		if !service.IsEnabledMqtt() {
			builder.EnableMQTT("192.168.88.242", "kaDUN6HF")
		}

		return builder
	})
	fmt.Println(err)

	var callback mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		if msg.Topic() != "megad/cmd" {
			var message megad2561.MegadMessage

			err := json.Unmarshal(msg.Payload(), &message)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(message)

			if message.Port == 0 {
				service.MqttClient.Publish("cmd", "31:2")
			}

			if message.Port == 17 { //nolint:gomnd
				service.MqttClient.Publish("cmd", "9:2")
			}

			if message.Port == 18 {
				service.MqttClient.Publish("cmd", "23:2")
			}
		}

		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())

		// 31 - workroom
	}

	service.MqttClient.Subscribe(callback)

	<-c
}
