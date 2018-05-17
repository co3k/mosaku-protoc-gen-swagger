# tetete

mbox.swagger.json の生成はこんな感じでいい？

```
$ protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=. mbox.proto
```

go コード生成 (RPC 用のコードは生成しない)

```
$ protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=. mbox.proto
```
