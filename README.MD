# 大的方向
## 关于项目分层

## 关于接口抽象



## 关于`gateway&grpc`接口数据结构设计
- `gateway`中的数据请求输入与输出结构是根据展示层（如`app/web`内容展示）的业务场景决定的。
- 内部的`grpc`接口服务输入与输出数据结构通常是通用性的设计，上层接口对其进行组合封装以满足上层不同的业务场景。
- 因此`gateway`与`grpc`的接口设计通常是不一样的。但在本示例项目中仅做演示，因此接口数据结构设计上有一定类似，请悉知。


## 管理`internal`目录的使用

# 其他一些细节
## `gateway json omitempty`
- 在`gateway`的接口设计中，没有对字段属性使用`json:"xxx,omitempty"`的配置，主要是考虑到展示层如果是`js`会频繁对字段进行
`typeof xxx != undefined`的判断会很繁琐，前后端协作效率很低。

## 关于字段硬编码


# 镜像编译、推送、验证

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