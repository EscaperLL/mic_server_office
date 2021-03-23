# Mic_srv_office Service

This is the Mic_srv_office service

Generated with

```
micro new --namespace=go.micro --type=service mic_srv_office
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.mic_srv_office
- Type: service
- Alias: mic_srv_office

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./mic_srv_office-service
```

Build a docker image
```
make docker
```