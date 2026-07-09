MODULE=github.com/I-Frostbyte/Hitch
PROTO_DIR=protos

# run server
run-server:
	go run cmd/server/main.go

build-docker:
	docker build -t hitch-server .
	docker run -p 40041:40041 hitch-server

build-docker-no-cache:
	docker build --no-cache -t hitch-server .
	docker run -p 40041:40041 hitch-server

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
