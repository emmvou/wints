# Wints

A web application to manage final internships at Polytech.

## Requirements

- A PostgreSQL database
- [Golang](https://golang.org/)
- [node.js](https://nodejs.org)

## Installation

In your `$GOPATH` (usually `$HOME/go`):
```shell
git clone https://github.com/emmvou/wints.git src/github.com/emmvou/wints
cd src/github.com/emmvou/wints/; npm ci
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

## Benchmarks, tests

go test -x -v -tags=integration -bench BenchmarkInternships -cpuprofile=cpu.prof
