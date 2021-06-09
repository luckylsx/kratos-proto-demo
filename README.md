# Kratos-proto-demo

生成代码

go generate ./...

运行：

cd cmd/kratos-proto-demo && go run . -conf ../../configs

如果运行失败

请查看configs 下配置文件 如数据库是否正确。

hello world
```api
curl 127.0.0.1:8000/helloworld/golang
```
emptypb request
```api
curl 127.0.0.1:8000/hello/sayEmpty
```

wrappers 类型 request
```api
curl 127.0.0.1:8000/hello/isDefault\?name=go\&age=0
```

response:
```api
{"name":"go", "age":0}
```

field_mask

request:
```api
curl -X POST -H "Content-Type:application/json" -d '{"task_id":1, "is_finished":true,"field_mask":"isFinished"}' 127.0.0.1:8000/hello/f
ieldMask
```

response
```api
{"fieldMask":["is_finished"]}
```

Any types request
```api
curl -X POST -H "Content-Type:application/json" -d '{"topic":"this is any type", "desc":{"@type":"type.googleapis.com/helloworld.v1.DescType", "value":"this is any type desc"}}' 127.0.0.1:8000/hello/any
```

response:
```api
{"topic":"this is any type", "desc":"\n\u0015this is any type desc"}
```

timestamp type request
```api
curl -X POST -H "Content-Type:application/json" -d '{"time_begin":"2021-06-08T15:15:30.069Z"}' 127.0.0.1:8000/hello/ts
```

response
```api
{"timestamp":"1623165330"}
```

Struct 类型：
```api
 curl -X POST -H "Content-Type:application/json" -d '{"json":{"name":"struct", "type":"int", "number":1}}' 127.0.0.1:8000/hello/struct
```

response

```api
{"detail":"{\"name\":\"struct\",\"number\":1,\"type\":\"int\"}"}%
```