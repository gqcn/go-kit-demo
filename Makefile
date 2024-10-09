
# compile proto files to pb go source files.
.PHONY: pb
pb:
	protoc --proto_path=user/manifest/proto \
	--go_out=paths=source_relative:user/api \
	--go-grpc_out=paths=source_relative:user/api \
	user/manifest/proto/api.proto

# compile services to images.
.PHONY: image
image:
	echo 1;

