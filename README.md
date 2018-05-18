# .proto を protoc-gen-swagger とか噛ましてどうにかいい感じにできないか模索する奴

とりあえずインストール

```
$ npm install
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

あと swagger-codegen を適当に入手してきて。

そんで make すればよい。

```
$ make
```
