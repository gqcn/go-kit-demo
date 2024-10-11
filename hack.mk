# compile services to binary.
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./temp/linux_amd64/main ./main.go

# compile proto files to pb go source files.
.PHONY: pb
pb:
	protoc --proto_path=manifest/proto \
	--go_out=paths=source_relative:api \
	--go-grpc_out=paths=source_relative:api \
	manifest/proto/*.proto


# compile services to images.
.PHONY: image
image: build
	$(eval _TAG  = $(shell git describe --dirty --always --tags --abbrev=8 --match 'v*' | sed 's/-/./2' | sed 's/-/./2'))
	$(eval _TAG  = $(if ${TAG},  ${TAG}, $(_TAG)))
	docker build --push --platform linux/amd64 -t $(DOCKER_NAME):${_TAG} -f manifest/docker/Dockerfile .

# compile helm to yaml.
.PHONY: yaml
yaml:
	helm template manifest/deploy > ./temp/deploy.yaml
