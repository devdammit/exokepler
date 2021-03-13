package devices_test

import (
	"github.com/stretchr/testify/assert"
	"keplerhub/pkg/devices"
	"testing"
)

func TestNewSwitchDevice(t *testing.T) {
	t.Parallel()

	t.Run("SwitchDevice Interface implements Device", func(t *testing.T) {
		t.Parallel()

		assert.Implements(t, (*devices.Device)(nil), new(devices.SwitchDevice))
	})

	testCases := []struct {
		name     string
		options  devices.SwitchDeviceOptions
		expected devices.SwitchDevice
	}{
		{
			name:    "Should return default",
			options: devices.SwitchDeviceOptions{},
			expected: devices.SwitchDevice{
				Pressed: false,
				Count:   0,
			},
		},
		{
			name: "Should reset counter when limit 1",
			options: devices.SwitchDeviceOptions{
				CounterLimit: 2,
			},
			expected: devices.SwitchDevice{
				Pressed: false,
				Count:   0,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			simpleSwitchDevice := devices.NewSwitchDevice(tc.options)

			assert.Equal(t, tc.expected.Pressed, simpleSwitchDevice.Pressed)
			assert.Equal(t, tc.expected.Count, simpleSwitchDevice.Count)
		})
	}
}

func TestSimpleSwitchDevice_Enable(t *testing.T) {
	t.Parallel()

	t.Run("Should pressed is true", func(t *testing.T) {
		options := devices.SwitchDeviceOptions{}
		simpleSwitchDevice := devices.NewSwitchDevice(options)

		assert.False(t, simpleSwitchDevice.Pressed)

		simpleSwitchDevice.Enable()

		assert.True(t, simpleSwitchDevice.Pressed)
	})

	t.Run("Should increment counter", func(t *testing.T) {
		options := devices.SwitchDeviceOptions{}

		simpleSwitchDevice := devices.NewSwitchDevice(options)

		assert.Equal(t, simpleSwitchDevice.Count, uint(0))

		simpleSwitchDevice.Enable()

		assert.Equal(t, simpleSwitchDevice.Count, uint(1))
	})

	t.Run("Should reset counter when out of limit pressed", func(t *testing.T) {
		options := devices.SwitchDeviceOptions{
			CounterLimit: 2,
		}

		simpleSwitchDevice := devices.NewSwitchDevice(options)

		assert.Equal(t, simpleSwitchDevice.Count, uint(0))

		simpleSwitchDevice.Enable()
		simpleSwitchDevice.Enable()

		assert.Equal(t, simpleSwitchDevice.Count, uint(2))

		simpleSwitchDevice.Enable()

		assert.Equal(t, simpleSwitchDevice.Count, uint(0))
	})
}

func TestSimpleSwitchDevice_Disable(t *testing.T) {
	t.Parallel()

	t.Run("Should set Pressed to falsy", func(t *testing.T) {
		options := devices.SwitchDeviceOptions{
			CounterLimit: 2,
		}

		simpleSwitchDevice := devices.NewSwitchDevice(options)

		simpleSwitchDevice.Enable()

		assert.True(t, simpleSwitchDevice.Pressed)

		simpleSwitchDevice.Disable()

		assert.False(t, simpleSwitchDevice.Pressed)
	})
}

func TestSimpleSwitchDevice_PowerToggle(t *testing.T) {
	t.Parallel()

	t.Run("Should toggle switch", func(t *testing.T) {
		options := devices.SwitchDeviceOptions{
			CounterLimit: 2,
		}

		simpleSwitchDevice := devices.NewSwitchDevice(options)

		assert.False(t, simpleSwitchDevice.Pressed)

		simpleSwitchDevice.PowerToggle()

		assert.True(t, simpleSwitchDevice.Pressed)
		assert.Equal(t, simpleSwitchDevice.Count, uint(1))

		simpleSwitchDevice.PowerToggle()

		assert.False(t, simpleSwitchDevice.Pressed)
		assert.Equal(t, simpleSwitchDevice.Count, uint(1))
	})
}
