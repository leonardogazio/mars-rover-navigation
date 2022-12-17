gen-proto:
	protoc proto/service.proto \
	--go_out=proto/ \
	--go-grpc_out=proto/ \
	--descriptor_set_out=proto/pb/descriptor.pb

tidy:
	go mod tidy

run:
	go run .

ui:
	grpcui --plaintext localhost:${HTTP_PORT}