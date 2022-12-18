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

ui:
	grpcui --plaintext localhost:${HTTP_PORT}