# tetete

mbox.swagger.json の生成はこんな感じでいい？

```
$ protoc --swagger_out="." mbox.proto
```

go コード生成 (RPC 用のコードは生成しない)

```
$ protoc --go_out=. mbox.proto
```
