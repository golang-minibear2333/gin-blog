# 环境准备
- docker（条件允许的话，建议配置docker加速器）
- docker-compose


## 配置代理

在执行前先设置好环境变量` GOPROXY`和` GO111MODULE`:
```shell
export GO111MODULE=on 

export GOPROXY=https://goproxy.io

或者

export GOPROXY=https://goproxy.cn 
```
对于1.13及以上版本，可直接如下这样
```go
go env -w GOPROXY=https://goproxy.cn,direct 
```
## 库安装
```go
go get -u github.com/gin-gonic/gin@v1.6.3
```
配置管理
```go
go get -u github.com/spf13/viper@v1.4.0
```
ORM 数据库连接工具
```go
go get -u github.com/jinzhu/gorm@v1.9.12
```
日志库: 单日志文件的最大占用空间、最大生存周期、允许保留的最多旧文件数（日志滚动），而我们使用这个库，主要是为了减免一些文件操作类的代码编写，把核心逻辑摆在日志标准化处理上
```go
go get -u gopkg.in/natefinch/lumberjack.v2
```
接口文档:使用` swag init`初始化
```go
go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
go get -u github.com/swaggo/gin-swagger@v1.2.0 
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
```
参数校验
```go
go get -u github.com/go-playground/validator/v10
```

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