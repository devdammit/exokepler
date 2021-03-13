package megadservice

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type ActionItemConfig struct {
	Type  string
	Value int16
}

type PortItemConfig struct {
	ID      int `yaml:"id"`
	Actions []ActionItemConfig
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
