gen-common-pb:
	rm -rf ./pb/common
	mkdir -p ./pb/common
	protoc --proto_path=./proto/common \
	--go_out=./pb/common \
	--go_opt=paths=source_relative \
 	--go-grpc_out=./pb/common \
 	--go-grpc_opt=paths=source_relative \
 	--go-grpc_opt=require_unimplemented_servers=false \
 	proto/common/*.proto

gen-user-service-go-pb:
	rm -rf ./pb/user-service-go
	mkdir -p ./pb/user-service-go
	protoc --proto_path=./proto/user-service-go \
	--go_out=./pb/user-service-go \
	--go_opt=paths=source_relative \
 	--go-grpc_out=./pb/user-service-go \
 	--go-grpc_opt=paths=source_relative \
 	--go-grpc_opt=require_unimplemented_servers=false \
 	proto/user-service-go/*.proto

gen-all: gen-user-service-go-pb gen-common-pb