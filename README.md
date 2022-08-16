## What is `protoc-gen-gin`
The `protoc-gen-gin` is a plugin of the Google protocol buffers compiler
[protoc](https://github.com/protocolbuffers/protobuf).
It reads protobuf service definitions and generates a [gin](https://github.com/gin-gonic/gin) server which
provide more RESTful HTTP API services. This server is generated according to the
[`google.api.http`](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L46)
annotations in your service definitions.

## Features supported
- Generating JSON API handlers.
- Method parameters in the request body.
- Method parameters in the request path.
- Method parameters in the query string.
- JWT supported and can generate OpenAPI v2 document.
- Enum fields in the path parameter (including repeated enum fields).
- Optionally emitting API definitions for
  [OpenAPI (Swagger) v2](https://swagger.io/docs/specification/2-0/basic-structure/).

## How to use
``` 
$ go install \
	google.golang.org/protobuf/cmd/protoc-gen-go@latest \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
	github.com/go-woo/protoc-gen-gin@latest

$ protoc --proto_path=. \
    --proto_path=./third_party \
    --go_out=paths=source_relative:. \
    --gin_out=paths=source_relative:. \
    ./v1/greeter.proto
```
## Generate OpenAPI v2 specification
```
$ protoc --proto_path=. \
    --proto_path=./third_party \
    --openapiv2_out . \
    --openapiv2_opt logtostderr=true \
    ./v1/greeter.proto
```
