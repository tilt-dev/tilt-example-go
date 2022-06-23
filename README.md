# tilt-example-go

[![Build Status](https://circleci.com/gh/tilt-dev/tilt-example-go/tree/master.svg?style=shield)](https://circleci.com/gh/tilt-dev/tilt-example-go)

An example project that demonstrates a live-updating Go server in Kubernetes. Read [doc](https://docs.tilt.dev/example_go.html).

- [0-base](0-base): The simplest way to start
- [1-measured](1-measured): Use `local_resource` to measure your deployment time
- [2-optimized](2-optimized): Compile Go binaries and copy them into Docker
- [3-recommended](3-recommended): Live-update compiled binaries

## Other Examples
- [tests-example](tests-example): an example of how to use Tilt to run your tests for you as you iterate

## License

Copyright 2022 Docker, Inc.

Licensed under [the Apache License, Version 2.0](LICENSE)
