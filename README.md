# Echo Protocol implemented with Go

[rfc862](https://datatracker.ietf.org/doc/html/rfc862)

"A very useful debugging and measurement tool is an echo service.  An
echo service simply sends back to the originating source any data it
receives."

## Run TCP server

```
ADDR=8090 go run . -opt tcp
```

## Run UDP server

```
ADDR=9090 go run . -opt udp
```

## Run client

### TCP

```
PROTOCOL=tcp ADDR=8090 go run . -opt client
```

### UDP

```
PROTOCOL=udp ADDR=9090 go run . -opt client
```

#### Echoed message

Start typing the message to be echoed back

## Run tests

```
go test -v ./...
```