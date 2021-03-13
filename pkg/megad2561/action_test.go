package megad2561_test

import (
	"github.com/stretchr/testify/assert"
	"keplerhub/pkg/megad2561"
	"testing"
)

func TestAction_Add(t *testing.T) {
	t.Parallel()

	t.Run("Should return value when added basic step", func(t *testing.T) {
		action := megad2561.Action{}

		action.Add(7, 1)

		assert.Equal(t, "7:1", action.GetValue())

		action.Add(30, 2)

		assert.Equal(t, "7:1;30:2", action.GetValue())
	})

	t.Run("Should return value with pause", func(t *testing.T) {
		action := megad2561.Action{}

		action.Add(7, 1).AddPause(100).Add(30, 2)
		assert.Equal(t, "7:1;p100;30:2", action.GetValue())
	})
}
