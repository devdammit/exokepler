package megad2561

type Service struct {
	config     Config
	MqttClient MqttClient
}

type BatchConfigureCallback func(builder BuilderConfig) BuilderConfig

type MegadService interface {
	GetConfig() Config
	HashID() bool
	IsEnabledMqtt() bool
	GetMQTTStatus() ConnectionStatusType
	BatchConfigure()
	FirmwareUpdate()
}

type MegadServiceOptions struct {
	HostIP
	Pwd
	SrvType
	SRV
	MQTTPwd
}

func NewMegadService(opts MegadServiceOptions) Service {
	s := Service{
		config: Config{
			hostIP: opts.HostIP,
			pwd:    opts.Pwd,
		},
	}

	s.config.syncFirstConfig()
	s.config.syncSecondConfig()

	if s.config.srvType == MQTT {
		s.enableMqtt()
	}

	return s
}

func (s *Service) BatchConfigure(cb BatchConfigureCallback) error {
	bc := newBuilderConfig(s.config)

	oldConfig := s.config
	updatedConfig := cb(bc)

	if s.config.SetAndUpdateConfig(updatedConfig) != nil {
		return ErrUpdateConfig
	}

	if oldConfig.srvType != updatedConfig.srvType {
		s.enableMqtt()
	}

	return nil
}

func (s *Service) GetConfig() Config {
	return s.config
}

func (s *Service) HasID() bool {
	return len(s.config.megadID) > 0
}

func (s *Service) IsEnabledMqtt() bool {
	return s.config.srvType == MQTT
}

func (s *Service) AddTrigger() {

}

/**
 * Private
**/

func (s *Service) enableMqtt() {
	mqttOpts := MqttClientOptions{Address: s.config.srv, ClientID: s.config.megadID, Password: s.config.mqttPassword}
	s.MqttClient = NewMqttClient(mqttOpts)
}
