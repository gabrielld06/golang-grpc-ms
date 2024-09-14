proto:
	@protoc \
		--proto_path=proto "proto/orders.proto" \
		--go_out=pb/orders \
		--go_opt=paths=source_relative \
		--go-grpc_out=pb/orders \
		--go-grpc_opt=paths=source_relative

orders:
	go run services/orders/*.go

kitchen:
	go run services/kitchen/*.go

.PHONY: proto orders kitchen
