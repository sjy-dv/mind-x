v0_init:
	- protoc --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./ --proto_path=./ ./v0/*.proto && protoc --go_out=./ --proto_path=./ ./v0/*.proto