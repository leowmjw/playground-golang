# BeeWax Service

This is the BeeWax service

Generated with

```
micro new github.com/leowmjw/playground-golang/MICRO/beeWax --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.beeWax
- Type: srv
- Alias: beeWax

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./beeWax-srv
```

Build a docker image
```
make docker
```