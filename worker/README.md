# Primelab CPU Worker

To build the filesystem of the worker VM, see details in agent

## Build Worker

```sh
go build
```

Static (for Alpine Linux):

```sh
go build -tags netgo -ldflags '-extldflags "-static"'
```

## Run

```sh
./worker
```

## Demo

### Healthcheck

```sh
» curl -i 127.0.0.1:8080/health
HTTP/1.1 200 OK
Content-Type: text/plain; charset=UTF-8
Date: Sat, 31 Jul 2021 20:02:45 GMT
Content-Length: 2
```
