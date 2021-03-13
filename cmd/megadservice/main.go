package main

import (
	"fmt"
	"keplerhub/internal/megadservice"
	"keplerhub/pkg/megad2561"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	var config megadservice.Config

	config.GetConf()

	options := megad2561.MegadServiceOptions{
		HostIP: "192.168.88.14",
		Pwd:    "sec",
	}
	service := megad2561.NewMegadService(options)

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

	megadservice.AddHandlers(service.MqttClient)

	<-c
}
