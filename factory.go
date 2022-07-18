package rpii2c

import (
	"github.com/d2r2/go-i2c"
	"github.com/d2r2/go-logger"
	corei2c "github.com/go-sensors/core/i2c"
	coreio "github.com/go-sensors/core/io"
	"github.com/pkg/errors"
)

type I2CPort struct {
	Address byte
	Bus     int
}

type wrapper struct {
	*i2c.I2C
}

func (w *wrapper) Read(p []byte) (n int, err error) {
	return w.ReadBytes(p)
}

func (w *wrapper) Write(p []byte) (n int, err error) {
	return w.WriteBytes(p)
}

func NewI2CPort(bus int, config *corei2c.I2CPortConfig) (*I2CPort, error) {
	i2cport := &I2CPort{
		Address: config.Address,
		Bus:     bus,
	}
	return i2cport, nil
}

func (p *I2CPort) Open() (coreio.Port, error) {
	i2c, err := i2c.NewI2C(p.Address, p.Bus)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open I2C address %v on bus %v", p.Address, p.Bus)
	}

	wrapper := &wrapper{
		i2c,
	}
	return wrapper, nil
}

func init() {
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
}
