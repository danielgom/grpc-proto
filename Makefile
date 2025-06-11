GO_MODULE := github.com/danielgom/grpc-proto

.PHONY: build protoc-go pipeline-lint pipeline-build clean-gateway protoc-go-gateway protoc-openapiv2-gateway

protoc-go:
	protoc --go_opt=module=$(GO_MODULE) --go_out=. \
  --go-grpc_opt=module=$(GO_MODULE) --go-grpc_out=. \
  ./proto/hello/*.proto ./proto/payment/*.proto ./proto/transaction/*.proto \
  ./proto/bank/*.proto ./proto/bank/type/*.proto \

build: protoc-go

pipeline-lint:
	sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

pipeline-build: pipeline-lint build

clean-gateway:
	rm -rf ./protogen/gateway
	mkdir -p ./protogen/gateway/go
	mkdir -p ./protogen/gateway/openapiv2

gateway-lint:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

protoc-go-gateway:
	protoc -I . \
	--plugin=protoc-gen-grpc_gateway=$(HOME)/go/bin/protoc-gen-grpc-gateway \
	--grpc_gateway_out ./protogen/gateway/go \
	--grpc_gateway_opt logtostderr=true \
	--grpc_gateway_opt paths=source_relative \
	--grpc_gateway_opt grpc_api_configuration=./grpc-gateway/config.yaml \
	--grpc_gateway_opt standalone=true \
	--grpc_gateway_opt generate_unbound_methods=true \
	./proto/hello/*.proto \
	./proto/bank/*.proto ./proto/bank/type/*.proto \

protoc-openapiv2-gateway:
	protoc -I . --openapiv2_out ./protogen/gateway/openapiv2 \
	--openapiv2_opt logtostderr=true \
	--openapiv2_opt output_format=yaml \
	--openapiv2_opt grpc_api_configuration=./grpc-gateway/config.yaml \
	--openapiv2_opt generate_unbound_methods=true \
	--openapiv2_opt allow_merge=true \
	--openapiv2_opt merge_file_name=merged \
	./proto/hello/*.proto \
    ./proto/bank/*.proto ./proto/bank/type/*.proto \

build-gateway: gateway-lint clean-gateway protoc-go-gateway protoc-openapiv2-gateway
