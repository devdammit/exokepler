package devices

type DeviceType int

//go:generate stringer -type=DeviceType -output=device_type_string.go
const (
	_ DeviceType = iota
	// Light is any lamp, chandelier, light bulbs.
	Light
	// Socket is connector, plug connector.
	Socket
	// Switch is light or something switch device.
	Switch
	// Thermostat is device with the possibility of temperature control.
	Thermostat
)

type Device interface {
	Enable()
	Disable()
	PowerToggle()
}
