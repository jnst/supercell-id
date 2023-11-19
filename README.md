# supercell-id

## API Testing

### curl

```sh
curl \
  --header 'Content-Type: application/json' \
  --data '{"email": "test@example.jp"}' \
  http://localhost:8080/authentication.v1.AuthenticationService/Register
```

### grpcurl

```sh
buf build -o proto/service.binpb
```

```sh
grpcurl -plaintext -protoset proto/service.binpb \
  -d '{"email": "test@example.jp"}' \
  localhost:8080 \
  authentication.v1.AuthenticationService/Register
```

```sh
buf curl -v --schema proto/service.binpb \
  --protocol grpc --http2-prior-knowledge \
  --data '{"email":"test@example.jp"}' \
  http://localhost:8080/authentication.v1.AuthenticationService/Register
```

## Notes for reproduction

```sh
go install github.com/bufbuild/buf/cmd/buf@v1.26.1
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.8.8
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@v1.11.1
```

```sh
which buf grpcurl protoc-gen-go protoc-gen-connect-go
```

```sh
cd proto
buf mod init
```
