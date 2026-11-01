# grpc-example

A minimal gRPC server and client written in Go, showing the basic workflow: define a service in Protocol Buffers, generate Go code, implement the server and call it from a client.

## Service

Defined in [`proto/test.proto`](proto/test.proto):

| RPC | Request | Response | Description |
|---|---|---|---|
| `Echo` | `ResponseRequest { msg }` | `ResponseRequest` | Returns the message it receives. |
| `GetUser` | `UserRequest { uuid }` | `UserResponse { name, age, email }` | Returns a sample user. |

## Project layout

```
proto/    Protocol Buffers service definition
gen/      Generated Go code (protoc output)
server/   gRPC server, listens on localhost:8085
client/   Client that calls Echo and GetUser
Makefile  Code generation commands
```

## Requirements

- Go
- `protoc` with the `protoc-gen-go` and `protoc-gen-go-grpc` plugins (only needed to regenerate code)

## Run

```bash
# regenerate the Go code from the .proto file (optional)
make create

# terminal 1
go run ./server

# terminal 2
go run ./client
```

## License

[MIT](LICENSE)
