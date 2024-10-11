# 构建二进制文件，通常用于本地测试。
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./temp/linux_amd64/main ./main.go

# 编译proto定义文件到go pb源码
.PHONY: pb
pb:
	protoc --proto_path=manifest/proto \
	--go_out=paths=source_relative:api \
	--go-grpc_out=paths=source_relative:api \
	manifest/proto/*.proto


# 编译二进制、镜像，并推送到镜像仓库。
# 这里由于是演示，推送到的是dockerhub镜像仓库。
.PHONY: image
image: build
	$(eval _TAG  = $(shell git describe --dirty --always --tags --abbrev=8 --match 'v*' | sed 's/-/./2' | sed 's/-/./2'))
	$(eval _TAG  = $(if ${TAG},  ${TAG}, $(_TAG)))
	docker build --push --platform linux/amd64 -t 937399771982.dkr.ecr.us-west-1.amazonaws.com/go-kit-demo-$(DOCKER_NAME):${_TAG} -f manifest/docker/Dockerfile .

# 编译helm生成部署的yaml文件（开发测试使用）
.PHONY: yaml
yaml:
	helm template manifest/deploy > ./temp/dev.yaml

# 编译helm生成部署的yaml文件（生产环境使用，values文件会不同）
.PHONY: yaml.prod
yaml.prod:
	helm template manifest/deploy -f manifest/deploy/values-prod.yaml > ./temp/$(DOCKER_NAME).yaml