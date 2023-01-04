PKGS ?= $(shell go list ./server/... ./rover/...)

gen-proto:
	protoc proto/service.proto \
	--go_out=proto/ \
	--go-grpc_out=proto/ \
	--descriptor_set_out=proto/pb/descriptor.pb

tidy:
	go mod tidy

run:
	go run . -api=$(api) -input=$(input)

build:
	go build -o ./bin/app

build-x:
	docker run --rm -v $(pwd):/go/src/app -w /go/src/app leonardogazio/go:1.18 gox -osarch="linux/amd64" -output="app"

ui:
	grpcui --plaintext localhost:${HTTP_PORT}

unit-test:
	go test -ldflags="-extldflags=-Wl,--allow-multiple-definition" -v ${PKGS}