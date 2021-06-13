# gin-blog

golang写的博客系统（Go 语言编程之旅第二章的练习）

本练习来自 [Go 语言编程之旅第二章的练习](https://golang2.eddycjy.com/posts/ch2/01-simple-server/)

本练习的作者源码 [go-programming-tour-book/blog-service](https://github.com/go-programming-tour-book/blog-service)

### 功能与目录

见 [CHANGELOG.md](CHANGELOG.md)

### 库

gin框架: Go 编写的一个 HTTP Web 框架,除了快以外，还具备小巧、精美且易用的特性，目前广受 Go 语言开发者的喜爱，是最流行的 HTTP Web 框架

```shell
go get -u github.com/gin-gonic/gin@v1.6.3
```

配置管理

```shell
go get -u github.com/spf13/viper@v1.4.0
```

ORM 数据库连接操作

```shell
go get -u github.com/jinzhu/gorm@v1.9.12
```

日志库: 单日志文件的最大占用空间、最大生存周期、允许保留的最多旧文件数（日志滚动），而我们使用这个库，主要是为了减免一些文件操作类的代码编写，把核心逻辑摆在日志标准化处理上

```shell
go get -u gopkg.in/natefinch/lumberjack.v2
```

接口文档
```shell
go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
go get -u github.com/swaggo/gin-swagger@v1.2.0 
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
```