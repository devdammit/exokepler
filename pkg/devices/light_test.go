package devices_test

import (
	"github.com/stretchr/testify/assert"
	"keplerhub/pkg/devices"
	"testing"
)

func TestLightDeviceInterface(t *testing.T) {
	t.Parallel()

	t.Run("LightDevice Interface implements Device", func(t *testing.T) {
		t.Parallel()

		assert.Implements(t, (*devices.Device)(nil), new(devices.LightDevice))
	})
}

func TestLightDevice_Enable(t *testing.T) {
	t.Parallel()

	t.Run("Should enable light device after call method Enable", func(t *testing.T) {
		options := devices.LightDeviceOptions{}
		lightDevice := devices.NewLightDevice(options)

		assert.Equal(t, lightDevice.PowerOn, false, "Light Device is disabled")

		lightDevice.Enable()

		assert.Equal(t, lightDevice.PowerOn, true, "Light Device is enabled")
	})
}

func TestLightDevice_Disable(t *testing.T) {
	t.Parallel()

	t.Run("Should disable light device after call method Disable", func(t *testing.T) {
		options := devices.LightDeviceOptions{}
		lightDevice := devices.NewLightDevice(options)

		lightDevice.Enable()
		lightDevice.Disable()
		assert.Equal(t, lightDevice.PowerOn, false, "Light Device is disabled")
	})
}

func TestLightDevice_PowerToggle(t *testing.T) {
	t.Parallel()

	t.Run("Should toggle power on light device after call method PowerToggle", func(t *testing.T) {
		options := devices.LightDeviceOptions{}
		lightDevice := devices.NewLightDevice(options)

		lightDevice.PowerToggle()

		assert.Equal(t, lightDevice.PowerOn, true, "Light Device is enabled")

		lightDevice.PowerToggle()

		assert.Equal(t, lightDevice.PowerOn, false, "Light Device is disabled")
	})
}

func TestLightDevice_MakeBrighter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		options  devices.LightDeviceOptions
		expected uint8
	}{
		{
			name:     "Should not change brightness value",
			options:  devices.LightDeviceOptions{},
			expected: 0,
		},
		{
			name: "Should change brightness to 1",
			options: devices.LightDeviceOptions{
				CanRangeBrightness: true,
			},
			expected: 1,
		},
		{
			name: "Should change brightness to 255 when over space value",
			options: devices.LightDeviceOptions{
				CanRangeBrightness: true,
				DefaultBrightness:  250,
				PrecisionRange:     20,
			},
			expected: 255,
		},
		{
			name: "Should change brightness to 240 when precision 20 and brightness 220",
			options: devices.LightDeviceOptions{
				CanRangeBrightness: true,
				DefaultBrightness:  220,
				PrecisionRange:     20,
			},
			expected: 240,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			lightDevice := devices.NewLightDevice(tc.options)

			lightDevice.MakeBrighter()

			assert.Equal(t, tc.expected, lightDevice.Brightness)
		})
	}
}

func TestLightDevice_MakeDarker(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		options  devices.LightDeviceOptions
		expected uint8
	}{
		{
			name:     "Should not change brightness value",
			options:  devices.LightDeviceOptions{},
			expected: 0,
		},
		{
			name: "Should change brightness to 1 when brightness 2 and default precision",
			options: devices.LightDeviceOptions{
				CanRangeBrightness: true,
				DefaultBrightness:  2,
			},
			expected: 1,
		},
		{
			name: "Should change brightness to 230 when over space value",
			options: devices.LightDeviceOptions{
				CanRangeBrightness: true,
				DefaultBrightness:  250,
				PrecisionRange:     20,
			},
			expected: 230,
		},
		{
			name: "Should change brightness to 0 when precision 100 and brightness 90",
			options: devices.LightDeviceOptions{
				CanRangeBrightness: true,
				DefaultBrightness:  90,
				PrecisionRange:     100,
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			lightDevice := devices.NewLightDevice(tc.options)

			lightDevice.MakeDarker()

			assert.Equal(t, tc.expected, lightDevice.Brightness)
		})
	}
}

func TestLightDeviceInstance(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		options  devices.LightDeviceOptions
		expected devices.LightDevice
	}{
		{
			name:    "Init simple Light",
			options: devices.LightDeviceOptions{},
			expected: devices.LightDevice{
				PowerOn:            false,
				Brightness:         0,
				Smooth:             0,
				CanRangeBrightness: false,
				CanSmooth:          false,
				RangeBrightness: devices.LightRangeBrightness{
					Min:       0,
					Max:       255,
					Precision: 0,
				},
			},
		},
		{
			name: "Init light with brightness",
			options: devices.LightDeviceOptions{
				CanRangeBrightness: true,
				MinRange:           10,
				MaxRange:           240,
				DefaultBrightness:  250,
				PrecisionRange:     10,
			},
			expected: devices.LightDevice{
				PowerOn:            false,
				Brightness:         240,
				Smooth:             0,
				CanRangeBrightness: true,
				CanSmooth:          false,
				RangeBrightness: devices.LightRangeBrightness{
					Min:       10,
					Max:       240,
					Precision: 10,
				},
			},
		},
		{
			name: "Should brightness is MinRage when minrange better brightness",
			options: devices.LightDeviceOptions{
				CanRangeBrightness: true,
				MinRange:           100,
				DefaultBrightness:  40,
			},
			expected: devices.LightDevice{
				Brightness:         100,
				CanRangeBrightness: true,
				RangeBrightness: devices.LightRangeBrightness{
					Min:       100,
					Max:       255,
					Precision: 1,
				},
			},
		},
		{
			name: "Init light with default precision",
			options: devices.LightDeviceOptions{
				CanRangeBrightness: true,
			},
			expected: devices.LightDevice{
				CanRangeBrightness: true,
				RangeBrightness: devices.LightRangeBrightness{
					Min:       0,
					Max:       255,
					Precision: 1,
				},
			},
		},
		{
			name: "Init with custom precision",
			options: devices.LightDeviceOptions{
				CanRangeBrightness: true,
				PrecisionRange:     10,
			},
			expected: devices.LightDevice{
				CanRangeBrightness: true,
				RangeBrightness: devices.LightRangeBrightness{
					Max:       255,
					Precision: 10,
				},
			},
		},
		{
			name: "Init with default power is true",
			options: devices.LightDeviceOptions{
				DefaultPower: true,
			},
			expected: devices.LightDevice{
				PowerOn: true,
				RangeBrightness: devices.LightRangeBrightness{
					Max: 255,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			lightDevice := devices.NewLightDevice(tc.options)

			assert.Equal(t, tc.expected, lightDevice)
		})
	}
}
