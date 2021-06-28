#  (2021-06-27)

### Features

* **jwt:** jwt api鉴权 ([f381a20](https://github.com/golang-minibear2333/gin-blog/commit/f381a209081a0b9b865589d88334609fceec6b7d)), closes [#12](https://github.com/golang-minibear2333/gin-blog/issues/12)
* **jwt:** jWT 权限校验模块 ([#13](https://github.com/golang-minibear2333/gin-blog/issues/13)) ([83236bc](https://github.com/golang-minibear2333/gin-blog/commit/83236bc06a13d3968f36b2c305f4f5f759f4b27b)), closes [#10](https://github.com/golang-minibear2333/gin-blog/issues/10)

### Bug Fixes

* **jwt:** 1. debug 模式跳过鉴权 2. 修复tag标签路由错误的bug ([c39befa](https://github.com/golang-minibear2333/gin-blog/commit/c39befa1cd6327a2bfc416531ea081c1fc4704d0))

#  (2021-06-20)


### Bug Fixes

* **swagger:** swagger测试tag时，无法调用任何结构，参入参数无法解析 ([b66be89](https://github.com/golang-minibear2333/gin-blog/commit/b66be898776144f6de3e290b6dcd52dc6b9a8031)), closes [#3](https://github.com/golang-minibear2333/gin-blog/issues/3)
* **swagger:** 增加POST UPDATE方法中的accept json注释 ([a4415f0](https://github.com/golang-minibear2333/gin-blog/commit/a4415f0229f22794afcca3b42790b197a13b56c8))
* **tag:** tag的启用状态为未启用，也就是0时，无法更新成功 ([2391510](https://github.com/golang-minibear2333/gin-blog/commit/23915107641528230b5085a27472414a1fd8edc1)), closes [#2](https://github.com/golang-minibear2333/gin-blog/issues/2)


### Features

* **tag:** tag相关 service和dao处理逻辑，及所有模块共有的字段回调 ([6ccfc9d](https://github.com/golang-minibear2333/gin-blog/commit/6ccfc9d40bf42aaae687c11902e6b06435f01cc2))
* **tag:** 更新tag router 完成接口 ([a344bae](https://github.com/golang-minibear2333/gin-blog/commit/a344baee6c8a8a7d376f2abe8b6a3b8a0d7cfb91))


#  (2021-06-13)


### Features

* gin 分页处理 ([572c950](https://github.com/golang-minibear2333/gin-blog/commit/572c9501ae291dafe26c03fdb8d4544d43a09567))
* mysql数据库gorm模块 ([a7ff87a](https://github.com/golang-minibear2333/gin-blog/commit/a7ff87a1174980f0d1beef9cb951689b4be1a6f1))
* string 类型转换 ([597741e](https://github.com/golang-minibear2333/gin-blog/commit/597741e38e39c370bf084d2e1b7021ba7e45304e))
* tcp 响应处理 ([56eac52](https://github.com/golang-minibear2333/gin-blog/commit/56eac52db57bdaf9a2de873cf30a459074f80fe6))
* 创建数据库脚本与model ([9c23f23](https://github.com/golang-minibear2333/gin-blog/commit/9c23f238a89208d80f5a7e4892091b6baad846fb))
* 初始化Http服务 ([2f64350](https://github.com/golang-minibear2333/gin-blog/commit/2f643506afac6af48d839fd306c15a4d5576a415))
* 新增router和controller ([9fc6162](https://github.com/golang-minibear2333/gin-blog/commit/9fc61621bde26e3181a0c64d87aa1f7c8f758a5c))
* 日志模块 ([786a134](https://github.com/golang-minibear2333/gin-blog/commit/786a134d0ef028d4823e821a9fd2dc5bcd193fad))
* 配置模块（配置文件，配置解析，全局变量初始化） ([a3a3112](https://github.com/golang-minibear2333/gin-blog/commit/a3a3112614a1f6438062b1f202b0d6901d69c745))
* 错误码标准化 ([3e79e5a](https://github.com/golang-minibear2333/gin-blog/commit/3e79e5a48fe1c7985f299916aa26b7b5c23814de))



#  (2021-06-12)


### Features

* 创建数据库脚本与model ([9c23f23](https://github.com/golang-minibear2333/gin-blog/commit/9c23f238a89208d80f5a7e4892091b6baad846fb))
* 初始化Http服务 ([2f64350](https://github.com/golang-minibear2333/gin-blog/commit/2f643506afac6af48d839fd306c15a4d5576a415))
* 新增router和controller ([9fc6162](https://github.com/golang-minibear2333/gin-blog/commit/9fc61621bde26e3181a0c64d87aa1f7c8f758a5c))



