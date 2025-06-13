gen-server:
	protoc \
	--proto_path=generated \
	--go_out=../go/pbuf --go_opt=paths=source_relative \
	--go-grpc_out=../go/pbuf --go-grpc_opt=paths=source_relative \
	generated/coba-frascati.proto

gen-client:
	protoc \
	--proto_path=generated \
	--go_out=../grpc-client-coba/pbuf --go_opt=paths=source_relative \
	--go-grpc_out=../grpc-client-coba/pbuf --go-grpc_opt=paths=source_relative \
	generated/coba-coba.proto