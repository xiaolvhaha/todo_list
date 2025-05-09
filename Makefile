API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: wire
wire:
	@wire

.PHONY: api
api:
	@protoc --proto_path=./api \
			--proto_path=./third_party \
	 		--go_out=paths=source_relative:./api \
 	       	--go-http_out=paths=source_relative:./api \
 	       	--go-grpc_out=paths=source_relative:./api \
		   	$(API_PROTO_FILES)