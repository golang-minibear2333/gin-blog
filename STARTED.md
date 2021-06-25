# 环境准备
- docker（条件允许的话，建议配置docker加速器）
- docker-compose


# 环境启动

## 数据库启动
项目已经配置相关的数据库文件，需要直接运行` docker-compose.yml`即可，启动命令如下

```shell
docker-compose up
```
注意启动这个文件需要确保3306端口没有被占用

## 项目启动
运行命令
```go
go run main.go 
```