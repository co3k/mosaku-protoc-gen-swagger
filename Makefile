PROTOC_OPTION=-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/

all:
	protoc -I. $(PROTOC_OPTION) --swagger_out=. shelves.proto
	protoc -I. $(PROTOC_OPTION) --go_out=. shelves.proto
	npx widdershins shelves.swagger.json -o shelves.md
	cd swagger-codegen-dist; swagger-codegen generate -i ../shelves.swagger.json -l go
