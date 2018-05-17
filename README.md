# .proto を protoc-gen-swagger とか噛ましてどうにかいい感じにできないか模索する奴

とりあえずインストール

```
$ npm install
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

mbox.swagger.json の生成はこんな感じでいい？

```
$ protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=. mbox.proto
```

go コード生成 (RPC 用のコードは生成しない)

```
$ protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=. mbox.proto
```

ドキュメント生成

```
$ npx widdershins mbox.swagger.json -o mbox.md
```

Swagger によるコードとドキュメントの生成。ぶっちゃけこれだけで充分か？

```
$ cd swagger-codegen-dist; swagger-codegen generate -i ../mbox.swagger.json -l go
```
