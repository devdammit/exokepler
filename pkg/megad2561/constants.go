package megad2561

type (
	PortType    int
	ActionValue int8
)

const (
	// DefaultPWD is default password for accessing the device.
	DefaultPwd = "sec"

	// MqttPort hard assigned in controller.
	MqttPort = "1883"

	// PwdMaxLength - controller-level restriction.
	PwdMaxLength = 3

	// ScriptNameMaxLength - controller-level restriction.
	ScriptNameMaxLength = 15

	// MegadIDMaxLength - Need some kind of restriction. It's not detected at the controller level.
	MegadIDMaxLength = 5
)

type (
	// HostIP is address controller in IP view.
	HostIP string

	// Pwd - password for accessing the device (maximum of 3 characters).
	Pwd string

	// GW - The gateway. It makes sense to specify only if the server is located outside the current IP network.
	// If omitted, the field displays the value 255.255.255.255.
	GW string

	// SrvType - Select protocol HTTP or MQTT.
	SrvType int

	// SRV - The IP address of the main server to which MegaD-2561 will send messages about triggered inputs.
	// After the IP address, you can specify the port. By default, 80.
	SRV string

	// Script - A script on the server that processes messages from the device and generates responses.
	// (maximum of ScriptNameMaxLength characters).
	Script string

	// MegadID is prefix path for access to device.
	// Default MegadID is sec. Then full path is 192.168.1.14/sec.
	MegadID string

	// MQTTPwd is password to mqtt server.
	MQTTPwd string

	// ConnectionStatusType.
	ConnectionStatusType int
)

//go:generate stringer -type=SrvType -output=srv_type_string.go
const (
	HTTP SrvType = iota
	MQTT
)

//go:generate stringer -type=PortType -output=port_type_string.go
const (
	_ PortType = iota
	IN
	OUT
	ADC
	DSen
	I2C
)

const (
	_ ConnectionStatusType = iota
	CONNECTED
	DISCONNECTED
)
