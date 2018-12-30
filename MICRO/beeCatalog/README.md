# BeeCatalog Service

This is the BeeCatalog service

Generated with

```
micro new github.com/leowmjw/playground-golang/MICRO/beeCatalog --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.beeCatalog
- Type: srv
- Alias: beeCatalog

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
./beeCatalog-srv
```

Build a docker image
```
make docker
```