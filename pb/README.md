

- `go get -u google.golang.org/grpc`
- `go get google.golang.org/grpc/cmd/protoc-gen-go-grpc`
- `protoc --go_out=.  --go-grpc_out=. pb/face_api.proto`