all:
	protoc -I. -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=. shelves.proto
	protoc -I. -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=. shelves.proto
	npx widdershins shelves.swagger.json -o shelves.md
	cd swagger-codegen-dist; swagger-codegen generate -i ../shelves.swagger.json -l go
