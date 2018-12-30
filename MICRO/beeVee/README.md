# BeeVee Service

This is the BeeVee service

Generated with

```
micro new github.com/leowmjw/playground-golang/MICRO/beeVee --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.beeVee
- Type: srv
- Alias: beeVee

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
./beeVee-srv
```

Build a docker image
```
make docker
```