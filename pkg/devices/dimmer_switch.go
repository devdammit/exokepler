package devices

type (
	DimmerSwitchDevice struct {
		Value uint8
	}
)

func NewDimmerSwitchDevice() DimmerSwitchDevice {
	return DimmerSwitchDevice{}
}

func (device *DimmerSwitchDevice) Enable() {
	device.Value = MaxRange
}

func (device *DimmerSwitchDevice) Disable() {
	device.Value = 0
}

func (device *DimmerSwitchDevice) PowerToggle() {
	if device.Value > 0 {
		device.Disable()
	} else {
		device.Enable()
	}
}

func (device *DimmerSwitchDevice) SetValue(v uint8) {
	device.Value = v
}
