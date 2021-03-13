package devices

// SwitchDevice is light or something switch device.
type SwitchDevice struct {
	Pressed      bool
	Count        uint
	counterLimit uint
}

type SwitchDeviceOptions struct {
	CounterLimit uint
}

// NewSwitchDevice is constructor for SwitchDevice.
func NewSwitchDevice(opts SwitchDeviceOptions) SwitchDevice {
	return SwitchDevice{
		counterLimit: opts.CounterLimit,
	}
}

// Enable implement method of interface Device.
func (device *SwitchDevice) Enable() {
	device.Pressed = true

	if device.counterLimit > 0 && device.Count >= device.counterLimit {
		device.Count = 0
	} else {
		device.Count++
	}
}

// Disable implements method of interface Device.
func (device *SwitchDevice) Disable() {
	device.Pressed = false
}

// PowerToggle implement interface Device.PowerToggle.
// Use SwitchDevice methods for control ur switch.
func (device *SwitchDevice) PowerToggle() {
	if device.Pressed {
		device.Disable()
	} else {
		device.Enable()
	}
}
