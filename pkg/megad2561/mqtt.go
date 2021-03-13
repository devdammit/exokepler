package megad2561

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	connection mqtt.Client
}

type MqttClientOptions struct {
	Address  SRV
	ClientID MegadID
	Password MQTTPwd
}

type MegadMessage struct {
	Port  int    `json:"port"`
	Value string `json:"value"`
}

func NewMqttClient(opts MqttClientOptions) MqttClient {
	address := fmt.Sprintf("tcp://%v", opts.Address)

	mqttOpts := mqtt.NewClientOptions().AddBroker(address).SetClientID("MegadGO")
	mqttOpts.SetKeepAlive(2 * time.Second)
	mqttOpts.SetPingTimeout(1 * time.Second)
	mqttOpts.SetPassword(string(opts.Password))
	mqttOpts.SetUsername(string(opts.ClientID))

	c := mqtt.NewClient(mqttOpts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return MqttClient{
		connection: c,
	}
}

func (mc *MqttClient) Subscribe(cb mqtt.MessageHandler) {
	if token := mc.connection.Subscribe("megad/#", 0, cb); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())

		os.Exit(1)
	}
}

func (mc *MqttClient) Publish(topic string, msg string) {
	client := mc.connection

	token := client.Publish(fmt.Sprintf("megad/%v", topic), 0, false, msg)
	token.Wait()
}
