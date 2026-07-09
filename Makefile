MODULE=github.com/I-Frosbyte/Hitch
PROTO_DIR=protos

# run server
run-server:
	go run cmd/server/main.go

# generate protos
generate-protos:
	@for file in $(PROTO_DIR)/*.proto; do \
		protoc \
			--go_out=. \
			--go_opt=module=$(MODULE) \
			--go-grpc_out=. \
			--go-grpc_opt=module=$(MODULE) \
			$$file; \
	done
