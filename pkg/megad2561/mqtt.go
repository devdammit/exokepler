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

type MegadPortInMessage struct {
	Port  int    `json:"port"`
	Mode  int8   `json:"m"`
	Value string `json:"value"`
	Click int8   `json:"click,omitempty"`
	Count int32  `json:"cnt"`
}

type MessageHandlerCallback func(msg MegadPortInMessage)

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
	fmt.Printf("publish %v:%v \n", topic, msg)
	token := client.Publish(fmt.Sprintf("megad/%v", topic), 0, false, msg)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
	}
}

func (mc *MqttClient) SendAction(action Action) {
	mc.Publish("cmd", action.value)
}

func (mc *MqttClient) AddHandler(in int, handler MessageHandlerCallback) error {
	topic := fmt.Sprintf("megad/%v", in)

	if token := mc.connection.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		var message MegadPortInMessage

		fmt.Println(string(msg.Payload()))

		err := json.Unmarshal(msg.Payload(), &message)
		if err != nil {
			fmt.Println(CantParseMqttMessage(msg))
		}

		if message.Port == in {
			handler(message)
		}
	}); token.Wait() && token.Error() != nil {
		return CantSubscribeToPort(in)
	}

	return nil
}
