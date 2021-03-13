package megad2561

import (
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	ErrUpdateConfig        = errors.New("cant update config")
	ErrMaxLengthID         = errors.New("device id maximum 5 length")
	ErrCantSubscribeToPort = errors.New("cant subscribe to port")
	ErrParseMqttMessage    = errors.New("cant parse mqtt message")
)

func CantSubscribeToPort(port int) error {
	return fmt.Errorf("%v %v", ErrCantSubscribeToPort, port)
}

func CantParseMqttMessage(msg mqtt.Message) error {
	return fmt.Errorf(
		"MQTT: %w. Original message '%v' and topic '%v'",
		ErrParseMqttMessage,
		msg.Payload(),
		msg.Topic(),
	)
}
