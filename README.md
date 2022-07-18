# go-sensors/rpii2c

Go library for providing [I2C] connectivity to sensors connected to Raspberry Pi computers (and compatible clones) using the [d2r2/go-i2c] package.

[i2c]: https://en.wikipedia.org/wiki/I%C2%B2C
[d2r2/go-i2c]: https://github.com/d2r2/go-i2c

## Quickstart

The I2C bus needs to be enabled in system configuration in order for this to work. See [d2r2/go-i2c Troubleshooting][troubleshooting] for more information.

Take a look at [rpi-sensor-exporter][rpi-sensor-exporter] for an example implementation that makes use of this library and several sensors.

[troubleshooting]: https://github.com/d2r2/go-i2c#troubleshooting
[rpi-sensor-exporter]: https://github.com/go-sensors/rpi-sensor-exporter

## Building

This software doesn't have any compiled assets.

## Code of Conduct

We are committed to fostering an open and welcoming environment. Please read our [code of conduct](CODE_OF_CONDUCT.md) before participating in or contributing to this project.

## Contributing

We welcome contributions and collaboration on this project. Please read our [contributor's guide](CONTRIBUTING.md) to understand how best to work with us.

## License and Authors

[![Daniel James logo](https://secure.gravatar.com/avatar/eaeac922b9f3cc9fd18cb9629b9e79f6.png?size=16) Daniel James](https://github.com/thzinc)

[![license](https://img.shields.io/github/license/go-sensors/rpii2c.svg)](https://github.com/go-sensors/rpii2c/blob/master/LICENSE)
[![GitHub contributors](https://img.shields.io/github/contributors/go-sensors/rpii2c.svg)](https://github.com/go-sensors/rpii2c/graphs/contributors)

This software is made available by Daniel James under the MIT license.
