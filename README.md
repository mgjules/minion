# minion

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](https://godoc.org/github.com/mgjules/minion)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=for-the-badge)](LICENSE)

A little minion (i.e microservice) that can be replicated to create more minions.

## Contents

- [minion](#minion)
  - [Contents](#contents)
  - [Requirements](#requirements)
  - [Architecture](#architecture)
  - [The why](#the-why)
  - [License](#license)
  - [Stability](#stability)

## Requirements

- [Go 1.20+](https://golang.org/doc/install)
- [Mage](https://github.com/magefile/mage) - replacement for Makefile in Go.
- [Golangci-lint](https://github.com/golangci/golangci-lint) - Fast Go linters runner.
- [Ginkgo](https://github.com/onsi/ginkgo) - Esspressive testing framework.
- [Docker](https://www.docker.com) - Containerization.
- [Docker-compose](https://docs.docker.com/compose/install/) - Orchestration of containers.

## Architecture

![Architecture](architecture.svg)

Inter-microservice communication is handled using gRPC instead of REST due to its higher performance, smaller payload size and tighter API contract.

Both Go services are instrumented using [OpenTelemetry](https://opentelemetry.io/) and use environment variables for configuration. Traces, metrics and logs are sent using gRPC to an OpenTelemetry collector, which batches and exports them to a Jaeger and Prometheus service.

## The why

Minion was created to simplify the creation of microservices. It is lightweight and opinionated.

Some of technologies used in this project are:

- [urfave/cli](https://github.com/urfave/cli) - Command line interface.
- [swag](https://github.com/swaggo/swag) - Generate REST API documentation.
- [alexliesenfeld/health](https://github.com/alexliesenfeld/health) - Simple and flexible health check library.
- [jsoniter](https://github.com/json-iterator/go) - High-performance drop-in replacement of `encoding/json`.
- [zap](https://github.com/uber-go/zap) - Blazing fast, structured, leveled logging in Go.
- [opentelemetry](https://github.com/open-telemetry/opentelemetry-go) - Open source distributed tracing and metrics.
- [watermill](https://github.com/ThreeDotsLabs/watermill) - Event messaging.
- [gin](https://github.com/gin-gonic/gin) - HTTP web framework.
- [resty](https://github.com/go-resty/resty) - HTTP client.

## License

Minion is Apache 2.0 licensed.

## Stability

This project follows [SemVer](http://semver.org/) strictly and is not yet `v1`.

Breaking changes might be introduced until `v1` is released.

This project follows the [Go Release Policy](https://golang.org/doc/devel/release.html#policy). Each major version of Go is supported until there are two newer major releases.
