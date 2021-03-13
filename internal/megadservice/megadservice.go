package megadservice

import (
	"fmt"
	"keplerhub/pkg/megad2561"
)

type ActionType int

//go:generate stringer -type=ActionType -output=action_type_string.go
const (
	_ ActionType = iota
	TOGGLE
	PAUSE
	REPEAT
)

func MapActions(actions []ActionItemConfig) *megad2561.Action {
	actionBuilder := megad2561.Action{}

	for _, action := range actions {
		switch action.Type {
		case TOGGLE.String():
			actionBuilder.Add(action.Value, 2)
		case PAUSE.String():
			actionBuilder.AddPause(action.Value)
		case REPEAT.String():
			actionBuilder.AddRepeat(int8(action.Value))
		}
	}

	return &actionBuilder
}

func AddHandlers(mqtt megad2561.MqttClient) {
	var config Config

	config.GetConf()

	for _, v := range config.Ports {
		err := mqtt.AddHandler(v.ID, MapActions(v.Actions))
		if err != nil {
			fmt.Println(err)
		}
	}

}
