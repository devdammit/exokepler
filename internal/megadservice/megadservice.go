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
	ON
	OFF
	PAUSE
	REPEAT
	GLOBAL
	DIMMER
	TDIMMER
)

func MapActions(actions []ActionItemConfig, builder *megad2561.Action) *megad2561.Action {
	for _, action := range actions {
		switch action.Type {
		case TOGGLE.String():
			builder.Add(action.Target, 2)
		case OFF.String():
			builder.Add(action.Target, 0)
		case ON.String():
			builder.Add(action.Target, 1)
		case DIMMER.String():
			builder.Add(action.Target, action.Value)
		case GLOBAL.String():
			builder.AddGlobal(action.Value)
		case REPEAT.String():
			builder.AddRepeat(action.Value)
		case TDIMMER.String():
			builder.AddShim(action.Target, action.Value)
		}
	}

	return builder
}

func PortMessageHandler(config PortItemConfig, client *megad2561.MqttClient) megad2561.MessageHandlerCallback {
	return func(msg megad2561.MegadPortInMessage) {
		actions := megad2561.Action{}
		modeType := ToModeType(config.Mode)
		handle := config.Handle

		if modeType == CLICK {
			if msg.Mode == 0 && msg.Click == 1 && len(handle.Single) > 0 {
				MapActions(handle.Single, &actions)
			} else if msg.Mode == 1 && len(handle.LongPress) > 0 {
				MapActions(handle.LongPress, &actions)
			} else if msg.Mode == 0 && msg.Click == 2 && len(handle.Double) > 0 {
				MapActions(handle.Double, &actions)
			}
		} else if modeType == PRESS {
			if msg.Mode == 2 && len(handle.Single) > 0 {
				MapActions(handle.Single, &actions)
			} else if msg.Mode == 1 && len(handle.LongPress) > 0 {
				MapActions(handle.LongPress, &actions)
			}
		} else if modeType == ANY {
			if len(handle.Single) > 0 {
				if msg.Mode == 2 || msg.Mode == 1 {
					MapActions(handle.Single, &actions)
				}
			}
		}
		fmt.Println(actions.GetValue())

		client.SendAction(actions)
	}
}

func AddHandlers(mqtt *megad2561.MqttClient) {
	var config Config

	config.GetConf()

	for _, v := range config.Ports {

		err := mqtt.AddHandler(v.ID, PortMessageHandler(v, mqtt))

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Added Handler %v \n", v.ID)

	}

}
