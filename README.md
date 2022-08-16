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

## Generate OpenAPI v2 specification
```
protoc -I . --openapiv2_out ./gen/openapiv2 \
--openapiv2_opt logtostderr=true \
your/service/v1/your_service.proto
```

``` 
swagger==={"info":{"title":"Greeter service","contact":{"name":"protoc-gen-gin example","url":"https://github.com/go-woo/protoc-gen-gin","email":"none@example.com"},"license":{"name":"MIT License","url":"https://github.com/go-woo/protoc-gen-gin/blob/main/LICENSE"},"version":"1.0"},"security_definitions":{"security":{"ApiKeyAuth":{"type":2,"name":"X-API-Key","in":2,"extensions":{"x-who-gen-token":"/login"}}}},"security":[{"security_requirement":{"ApiKeyAuth":{}}}]}
```