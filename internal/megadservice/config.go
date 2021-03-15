package megadservice

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type ModeType int

//go:generate stringer -type=ModeType -output=mode_type_string.go
const (
	_ ModeType = iota
	CLICK
	PRESS
	ANY
)

type ActionItemConfig struct {
	Type   string
	Target int16
	Value  int16
}

type PortModes struct {
	Single    []ActionItemConfig `yaml:"single,omitempty"`
	Double    []ActionItemConfig `yaml:"double,omitempty"`
	LongPress []ActionItemConfig `yaml:"long_press,omitempty"`
}

type PortItemConfig struct {
	ID     int    `yaml:"id"`
	Mode   string `yaml:"mode"`
	Handle PortModes
}

type Config struct {
	Settings struct {
		HostIP       string `yaml:"host_ip"`
		Pwd          string `yaml:"pwd"`
		Srv          string `yaml:"srv"`
		SrvType      string `yaml:"srv_type"`
		MqttPassword string `yaml:"mqtt_password"`
	}

	Ports []PortItemConfig
}

func (c *Config) GetConf() *Config {
	file, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Printf("Cant get config #%v", err)
	}

	err = yaml.Unmarshal(file, c)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func ToModeType(source string) ModeType {
	var modeType ModeType

	switch source {
	case CLICK.String():
		modeType = CLICK
	case PRESS.String():
		modeType = PRESS
	case ANY.String():
		modeType = ANY
	default:
		modeType = CLICK
	}

	return modeType
}
