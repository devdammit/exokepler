package megad2561

import "fmt"

type BuilderConfig struct {
	srv          SRV
	srvType      SrvType
	megadID      MegadID
	mqttPassword MQTTPwd
}

func newBuilderConfig(c Config) BuilderConfig {
	bc := BuilderConfig{}

	bc.srv = c.srv
	bc.srvType = c.srvType
	bc.megadID = c.megadID
	bc.mqttPassword = c.mqttPassword

	return bc
}

// EnableMQTT set MQTT params to enabled.
func (bc *BuilderConfig) EnableMQTT(srv SRV, password MQTTPwd) {
	bc.srv = SRV(fmt.Sprintf("%v:%v", srv, MqttPort))
	bc.srvType = MQTT
	bc.mqttPassword = password
}

// SetID set id megad controller.
func (bc *BuilderConfig) SetID(id MegadID) error {
	if len(id) > MegadIDMaxLength {
		return ErrMaxLengthID
	}

	bc.megadID = id

	return nil
}

// DisableMQTT set MQTT params to disabled and set srv to default.
func (bc *BuilderConfig) DisableMQTT() {
	bc.srvType = HTTP
	bc.srv = ""
}
