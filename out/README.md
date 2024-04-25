# 阿里云OSS



## 启动

> 运行 oss.sh / oss.exe 会自动读取config.json的内容
>
> shell 会根据json内配置的端口启动http服务

## 注意事项

- 使用前请确保已获取config.json对应的参数
- 修改config.json参数时请勿修改任何的key
- sms.sh 只能在amd64下的linux环境运行
- 请确保config.json 在oss.sh / oss.exe同级目录下
- 配置http端口时请确保port的值为string类型

### API

> 成功启动通过 ip:port/uri 访问 API

| uri   | method | param |Content-Type | response | desc   |
|-------| ------ | ----- | -------- |----------|--------|
| upload| POST |file:*multipart.FileHeader| multipart/form-data| message:string| 上传文件   |
| single|DELETE| path:string| application/json| message:string| 删除单文件  |
| multiple| DELETE|paths:[]string|application/json| message:string| 删除多个文件 |


> 文件只能单张上传


