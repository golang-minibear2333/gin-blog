# 环境准备

- docker（条件允许的话，建议配置docker加速器）
- docker-compose up -d

## 配置代理

在执行前先设置好环境变量` GOPROXY`和` GO111MODULE`:

```shell
export GO111MODULE=on 

export GOPROXY=https://goproxy.io

或者

export GOPROXY=https://goproxy.cn 
```

对于1.13及以上版本，可直接如下这样

```shell
go env -w GOPROXY=https://goproxy.cn,direct 
```

## 库安装

```shell
go mod tidy
go mod vendor
```

* 使用`tidy`命令从网络上下载库
* `vendor`命令把库缓存到本项目中
* `.gitignore` 已配置，不提交`vendor`文件夹

库已设置到`go mod`中，不需要执行下面的命令，这里只作为库的作用提示

```shell
go get -u github.com/gin-gonic/gin@v1.6.3
```

配置管理

```shell
go get -u github.com/spf13/viper@v1.4.0
```

ORM 数据库连接工具

```shell
go get -u github.com/jinzhu/gorm@v1.9.12
```

日志库: 单日志文件的最大占用空间、最大生存周期、允许保留的最多旧文件数（日志滚动），而我们使用这个库，主要是为了减免一些文件操作类的代码编写，把核心逻辑摆在日志标准化处理上

```shell
go get -u gopkg.in/natefinch/lumberjack.v2
```

接口文档:使用` swag init`初始化

```shell
go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
go get -u github.com/swaggo/gin-swagger@v1.2.0 
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
```

参数校验

```shell
go get -u github.com/go-playground/validator/v10
```

panic告警邮件功能

```shell
go get -u gopkg.in/gomail.v2
```

服务限流

```shell
go get -u github.com/juju/ratelimit@v1.0.1
```

链路追踪

```shell
go get -u github.com/opentracing/opentracing-go@v1.0.0
go get -u github.com/uber/jaeger-client-go@v2.22.1
```

* 接入方法 `docker-compose up -d jaeger` 启动追踪服务并访问`http://localhost:16686/`转到`jaeger WEB UI`
* 启动`gin-blog`程序访问`localhost:8000/swagger/index.html`，随意访问一个接口
* 可以在追踪服务界面看到效果

SQL 追踪

```shell
go get -u github.com/eddycjy/opentracing-gorm
```
# 配置文件

## 邮件配置

QQ 邮件的 SMTP，这个只需要在”QQ 邮箱-设置-账户-POP3/IMAP/SMTP/Exchange/CardDAV/CalDAV 服务“选项中将”POP3/SMTP 服务“和”IMAP/SMTP 服务“开启，然后根据所获取的 SMTP 账户密码进行设置即可，另外 SSL 是默认开启的。

```yaml
Email:
  Host: smtp.qq.com
  Port: 465
  UserName: xxxx@qq.com
  Password: xxxxxxxx
  IsSSL: true
  From: xxxx@qq.com
  To:
    - xxxx@qq.com
```

另外需要特别提醒的一点是，我们所填写的 SMTP Server 的 HOST 端口号是 465，而常用的另外一类还有 25 端口号 ，但我强烈不建议使用 25，你应当切换为 465，因为 25 端口号在云服务厂商上是一个经常被默认封禁的端口号，并且不可解封，使用 25 端口，你很有可能会遇到部署进云服务环境后告警邮件无法正常发送出去的问题。

# 环境启动

## 数据库启动

项目已经配置相关的数据库文件，需要直接运行` docker-compose.yml`即可，启动命令如下

```shell
docker-compose up -d
```

注意启动这个文件需要确保3306端口没有被占用

## 项目启动

运行命令

```shell
go run main.go 
```