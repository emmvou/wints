# Wints

[![CI](https://github.com/emmvou/wints/actions/workflows/go.yml/badge.svg)](https://github.com/emmvou/wints/actions/workflows/go.yml)

A web application to manage final internships at Polytech.  
See [this repo](https://github.com/fhermeni/wints)

## Requirements

- A PostgreSQL database
- [Golang](https://golang.org/)
- [node.js](https://nodejs.org)

## Installation

In your `$GOPATH` (usually `$HOME/go`):
```shell
git clone https://github.com/emmvou/wints.git src/github.com/emmvou/wints
cd src/github.com/emmvou/wints/; npm ci; gulp assets
go install github.com/emmvou/wints
```

## Usage
In the wints installation directory (`$GOPATH/src/github.com/emmvou/wints`):
```shell
wints
```

### Parameters

```Shell
  -conf string  
        Wints configuration file (default "wints.conf")  
  -fake-mailer  
        Do not send emails. Print them out stdout  
  -install-db  
        install the database  
  -new-root string  
        Invite a root user  
```

## Development environment

In a development environment, use the following commands:
```shell
gulp assets watch
```

## Benchmarks, tests

```shell
go test -x -v -tags=integration -bench BenchmarkInternships -cpuprofile=cpu.prof -db-url $CONNECTION_STRING
```
Replace `$CONNECTION_STRING` with the string used in `wints.conf`.
