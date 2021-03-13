package megad2561

import (
	"encoding/json"
	"fmt"
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

func (mc *MqttClient) Publish(topic string, msg string) {
	client := mc.connection

	token := client.Publish(fmt.Sprintf("megad/%v", topic), 0, false, msg)
	token.Wait()
}

func (mc *MqttClient) AddHandler(in int, action *Action) error {
	topic := fmt.Sprintf("megad/%v", in)
	if token := mc.connection.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		var message MegadMessage

		err := json.Unmarshal(msg.Payload(), &message)
		if err != nil {
			fmt.Println(CantParseMqttMessage(msg))
		}

		if message.Port == in {
			mc.Publish("cmd", action.value)
		}
	}); token.Wait() && token.Error() != nil {
		return CantSubscribeToPort(in)
	}

	return nil
}
