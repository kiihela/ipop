## ipop
This project allows you to abstract away to `struct` types used in `github.com/gobuffalo/pop`.

## Motivation
This allows you to abstract away the real [`pop`](https://github.com/gobuffalo/pop) package with your own implementation, whether for some new backend or just for mocking in your unit tests.

## Build status

[![Build Status](https://travis-ci.org/dnnrly/ipop.svg?branch=master)](https://travis-ci.org/dnnrly/ipop)
[![codecov](https://codecov.io/gh/dnnrly/ipop/branch/master/graph/badge.svg)](https://codecov.io/gh/dnnrly/ipop)
[![](https://godoc.org/github.com/dnnrly/ipop?status.svg)](http://godoc.org/github.com/dnnrly/ipop)
[![](https://goreportcard.com/badge/github.com/dnnrly/ipop)](https://goreportcard.com/report/github.com/dnnrly/ipop)

## Installation
```
go get github.com/dnnrly/ipop
```

Or if you're using modules then just import it in your code.

## API Reference
This aims to be API compatible with [`pop`](https://github.com/gobuffalo/pop).

## Tests
Run tests by using the command:
```bash
make test
```

## How to use?
See the godoc examples.

## Contribute
Please see the [contributing guideline](https://github.com/dnnrly/ipop/blob/master/CONTRIBUTING.md).

## Credits
This is basically just wrapping up [`pop`](https://github.com/gobuffalo/pop). All the praise goes to them - and then some. 

## License

Apache2 © [Pascal Dennerly](dnnrly@gmail.com)
