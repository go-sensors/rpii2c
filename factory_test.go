package rpii2c_test

import (
	"testing"

	corei2c "github.com/go-sensors/core/i2c"
	"github.com/go-sensors/rpii2c"
	"github.com/stretchr/testify/assert"
)

func Test_NewI2CPort_returns_configured_I2C_port(t *testing.T) {
	// Arrange
	config := corei2c.I2CPortConfig{
		Address: 0x77,
	}

	// Act
	actual, err := rpii2c.NewI2CPort(1, &config)

	// Assert
	assert.NoError(t, err)
	assert.EqualValues(t, 1, actual.Bus)
	assert.EqualValues(t, 0x77, actual.Address)
}
