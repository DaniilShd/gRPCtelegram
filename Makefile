PHONY: generate-structs
generate-structs:
	mkdir -p pkg/user_v1
	protoc --go_out=pkg/user_v1 --go_opt=paths=source_relative \
	api/user_v1/service.proto

PHONY: generate
generate: 
	mkdir -p pkg/user_v1
	protoc --go_out=pkg/user_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative api/user_v1/service_grpc.proto
