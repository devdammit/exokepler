package devices

const (
	MaxRange = 255
)

type LightRangeBrightness struct {
	Min       uint8
	Max       uint8
	Precision uint8
}

// LightDevice is any lamp, chandelier, light bulbs.
type LightDevice struct {
	PowerOn            bool
	Brightness         uint8
	Smooth             uint8
	CanRangeBrightness bool
	CanSmooth          bool
	RangeBrightness    LightRangeBrightness
}

// LightDeviceOptions is options for create LightDevice.
type LightDeviceOptions struct {
	CanRangeBrightness bool
	CanSmooth          bool
	DefaultPower       bool
	Smooth             uint8
	MinRange           uint8
	MaxRange           uint8
	PrecisionRange     uint8
	DefaultBrightness  uint8
}

// NewLightDevice is constructor for LightDevice.
func NewLightDevice(opts LightDeviceOptions) LightDevice {
	defaultBrightness := opts.DefaultBrightness

	if opts.MaxRange != 0 && defaultBrightness > opts.MaxRange {
		defaultBrightness = opts.MaxRange
	}

	if defaultBrightness < opts.MinRange {
		defaultBrightness = opts.MinRange
	}

	if !opts.CanRangeBrightness {
		defaultBrightness = 0
	}

	ld := LightDevice{
		CanRangeBrightness: opts.CanRangeBrightness,
		Brightness:         defaultBrightness,
	}

	if opts.PrecisionRange == 0 && opts.CanRangeBrightness {
		ld.RangeBrightness.Precision = 1
	} else {
		ld.RangeBrightness.Precision = opts.PrecisionRange
	}

	if opts.MaxRange == 0 {
		ld.RangeBrightness.Max = MaxRange
	} else {
		ld.RangeBrightness.Max = opts.MaxRange
	}

	ld.PowerOn = opts.DefaultPower
	ld.RangeBrightness.Min = opts.MinRange
	ld.Smooth = opts.Smooth
	ld.CanSmooth = opts.CanSmooth

	return ld
}

// Enable implement method of interface Device.
func (device *LightDevice) Enable() {
	device.PowerOn = true
}

// Disable implements method of interface Device.
func (device *LightDevice) Disable() {
	device.PowerOn = false
}

// PowerToggle implement interface Device.PowerToggle.
// Use LightDevice methods for control brightness.
func (device *LightDevice) PowerToggle() {
	if device.PowerOn {
		device.Disable()
	} else {
		device.Enable()
	}
}

// MakeBrighter makes the brightness brighter.
func (device *LightDevice) MakeBrighter() uint8 {
	if device.CanRangeBrightness {
		if (MaxRange - device.Brightness) < device.RangeBrightness.Precision {
			device.Brightness = device.RangeBrightness.Max
		} else {
			device.Brightness += device.RangeBrightness.Precision
		}

		return device.Brightness
	}

	return device.Brightness
}

// MakeDarker makes the brightness darker.
func (device *LightDevice) MakeDarker() uint8 {
	if device.CanRangeBrightness {
		if (MaxRange - device.RangeBrightness.Precision) < (MaxRange - device.Brightness) {
			device.Brightness = 0
		} else {
			device.Brightness -= device.RangeBrightness.Precision
		}
	}

	return device.Brightness
}

// ChangeBrightness use that for change value. TODO: Add limit logic.
func (device *LightDevice) ChangeBrightness(v uint8) {
	device.Brightness = v
}
