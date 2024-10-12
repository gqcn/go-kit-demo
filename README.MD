本项目用于演示使用`go-kit`开发`grpc&http`服务，并使用到了`mongodb`做存储，`helm`作为部署管理工具，部署环境为`kubernetes`集群。

# 大的方向
## 关于项目分层
使用了`go-kit`官方的结构：
- `transport`：协议转换层。比如前端`json`到后端`endpoint`的数据结构转换，路由转发绑定。
- `endpoint`：具体的接口处理层。前端一个接口对应后端一个`endpoint`。
- `service`：业务逻辑抽象层。将`endpoint`中可复用的业务逻辑抽象沉淀到`service`层，以便各个接口复用，内部逻辑复用。
- `repository`：数据存储抽象层。将可复用的数据库操作抽象沉淀到这一层。但在大部分中小型项目中，对数据库操作的抽象沉淀是比较复杂的艺术，因此
并不是所有对数据库的操作都需要封装到这里维护。无法复用的数据库操作也允许在更上层去执行。

此外根据业务项目需要增加了：
- `api`：接口定义层。定义服务间、前后端交互的输入输出数据结构，例如`grpc/http`的输入输出定义。
- `model`：数据实体数据结构。该层的数据结构与数据库的数据集合字段结构应当一一对应，所以通常是自动化同步维护，人工维护可能会引发差异。此外，
业务逻辑相关的数据结构应当闭环到`endpoint`和`service`内部，不要定义到`model`中。
- `manifest`：项目交付清单。非二进制编译打包的资源文件，例如镜像编译文件、部署文件、

## 关于逻辑抽象
项目中有对业务逻辑和数据操作进行抽象的工作，抽象通常是需要根据已有业务场景实践、存在重复逻辑的基础上，例如`endpoint`在早期并没有重复逻辑，
不需要根据经验过早地去做抽象，等后续业务场景丰富后再根据具体场景实践去做抽象，将重复逻辑沉淀到`service`层。对于`repository`也是同理，但对于
`repository`来讲，它早期需要维护数据库集合的信息以及对象化，而不是暴露数据库连接对象让业务逻辑四处使用。

## 关于接口化设计
接口化通常有三种常见的使用场景：
1. 通常软件开发是先定义接口，再去完成具体的实现，是一种自顶向下的设计。
2. 同时，通过接口化也可以将各个业务模块进行解耦，只对外暴露需要的接口信息，而不是通过实例化对象完整暴露能力。
3. 此外，接口化可以实现很好的数据`mock`，针对单元测试来说比较有用。

针对`endpoint/servce/repository`内部会有很多具体的实现，但要不要再重复定义一个接口？如果咱们的习惯是先做实现再定义接口，个人建议这样的
场景下是没有太大必要再重复去定义一个接口的，会增加额外的开发维护成本。这种场景下可以使用可见性封装来达到目的即可，成本会更低。


## 关于`gateway&grpc`接口数据结构设计
- `gateway`中的数据请求输入与输出结构是根据展示层（如`app/web`内容展示）的业务场景决定的。
- 内部的`grpc`接口服务输入与输出数据结构通常是通用性的设计，上层接口对其进行组合封装以满足上层不同的业务场景。
- 因此`gateway`与`grpc`的接口设计通常是不一样的。但在本示例项目中仅做演示，因此接口数据结构设计上有一定类似，请悉知。

## 管理`internal`目录的使用
在大仓代码管理中，充分使用`Golang`语言提供的`internal`目录特性可以防止服务间意外依赖，也可以让现代`IDE`更智能去自动识别`import`依赖引入。
特别是每个服务内部会有一些相同包名的场景下，只公开需要公开的，不需要公开的都放到`internal`目录中内部使用，也能提高代码管理的简洁性。

# 其他一些细节

## 接口兼容性
使用`api`版本以保证接口向下兼容性，存在不兼容接口变更需要使用新的接口定义。

## `gateway json omitempty`
- 在`gateway`的接口设计中，没有对字段属性使用`json:"xxx,omitempty"`的配置，主要是考虑到展示层(特别是`web`)如果是`js`会频繁对字段进行
`typeof xxx != undefined`的判断会很繁琐，前后端协作效率比较低。

## 关于字段硬编码
曾经因为字段硬编码引发了生产环境的事故后，针对字段的操作通常都通过常量或者固定的对象属性来操作，这些字段可以自动化生成，不需要人工维护。

# 镜像编译、服务部署、接口验证
使用了`Makefile`来封装常用的工程管理命令。

## 镜像编译
通过`make image`编译并推送镜像。由于是演示项目，没有包含任何业务信息，所以目前使用的是镜像仓库是`dockerhub`。

## 服务部署
- 项目通过`helm`管理部署文件。
- 通过`make yaml`编译生成开发测试的`yaml`文件读取的是`values.yaml`文件。
- 通过`make yaml.prod`编译生成生产环境的`yaml`文件读取的是`values-prod.yaml`文件，该文件存放在本地未上传到仓库。

## 镜像验证
gateway服务：
```bash
docker run --platform linux/amd64 \
-e SERVICE_USER=user-service \
-e SERVER_ADDRESS=:80 \
loads/go-kit-demo-gateway:7c1e2e8a-dirty 
```

grpc服务：
```bash
docker run --platform linux/amd64 \
-e MONGODB_URI=mongodb://localhost:27017 \
-e SERVER_ADDRESS=:8000 \
loads/go-kit-demo-user-service:7c1e2e8a-dirty 
```

# 接口测试
部署后通过以下指令进入到`gateway`发送请求测试。
```bash
kubectl exec -it -n john gateway-64bd444b9b-xrt9s -- bash
```

## 用户创建接口 - Create
正确的请求：
```bash
curl -XPOST -d'{
  "user":{
    "name":"test_user",
    "age":18,
    "gender":1,
    "location":{
      "type":"Point",
      "coordinates":[104.06055237921923,30.5967693328857]
    },
    "match_gender": 2,
    "match_min_age": 18,
    "match_max_age": 30
  }
}' localhost:8080/create
```

错误的请求：
```bash
curl -XPOST -d'{
  "user":{
    "name":"",
    "age":18,
    "gender":1,
    "location":{
      "type":"Point",
      "coordinates":[104.06055237921923,30.5967693328857]
    },
    "match_gender": 2,
    "match_min_age": 18,
    "match_max_age": 30
  }
}' localhost:8080/create
```

## 用户查询接口 - Search

```bash
curl -XPOST -d'{
    "gender":1,
    "min_age": 18,
    "max_age": 30
}' localhost:8080/search
```