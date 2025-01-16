GO_MODULE := github.com/danielgom/grpc-proto

.PHONY: build protoc-go pipeline-lint pipeline-build
protoc-go:
	protoc --go_opt=module=$(GO_MODULE) --go_out=. \
  --go-grpc_opt=module=$(GO_MODULE) --go-grpc_out=. \
  ./proto/hello/*.proto ./proto/payment/*.proto ./proto/transaction/*.proto \
  ./proto/bank/*.proto ./proto/bank/type/*.proto \
  ./proto/resiliency/*.proto

build: protoc-go

pipeline-lint:
	sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

pipeline-build: pipeline-lint build
