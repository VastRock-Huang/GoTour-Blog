# GoTour-BlogService
* Ref: 《Go应用编程之旅：一起用Go做项目》第二章
## 实现功能
该项目主要使用 Go 完成了一个博客的后端开发, 实现功能比较简单, **核心功能为标签(tag)和文章(article)的增删查改操作**, 并实现了其它后端程序的基础功能如日志记录、访问控制、配置管理等等.
* 主要内容参考了《Go应用编程之旅：一起用Go做项目》一书的第二章, 源码即书作者的 GitHub 源码 [go-programming-tour-book/blog-service](https://github.com/go-programming-tour-book/blog-service)

## 目录结构
该项目的目录结构如下:  
```
gotour-blog_service
|-- configs
|-- docs
|-- global
|-- internal
    |-- dao
    |-- middleware
    |-- model
    |-- routers
    |-- service
|-- pkg
    |-- app
    |-- convert
    |-- email
    |-- errcode
    |-- limiter
    |-- logger
    |-- setting
    |-- tracer
    |-- upload
    |-- util
|-- storage
    |-- logs
    |-- uploads
```
* `configs`: 配置文件目录. 主要为配置文件 `config.yaml`.
* `docs`: 文档目录. 主要为由 `Swagger` 库生成的接口文档相关文件. 此外 `./sql/blog.sql` 中记录了项目中使用的 SQL 语句.
* `global`: 全局变量. 即整个程序中全局唯一的变量, 可以全局共同使用, 包括配置结构体,日志对象,数据库对象和链路追踪对象.
* `internal`: 内部模块. 主要为和业务功能联系紧密的代码模块.
    * `dao`: 数据库访问层(Database Access Object): 介于上层数据和下层数据库模型操作之间的一层.
    * `middleware`: Gin 中间件.
    * `model`: 模型层, 即数据库中数据表的模型对象, 为数据操作最底层.
    * `routers`: 路由操作相关的逻辑.
    * `service`: 项目核心业务逻辑, 位于数据操作的最上层.
* `pkg`: 项目相关的模块包. 与业务功能相关性不大, 基本为后端项目的通用模块包.
    * `app`: 应用模块. 主要与与请求的参数绑定,报错,令牌验证以及响应相关.
    * `convert`: 字符串的类型转换转换模块.
    * `email`: 发送邮件模块.
    * `errcode`: 项目错误码标准化模块.
    * `limiter`: 接口限流模块.
    * `logger`: 日志模块.
    * `setting`: 配置模块.
    * `tracer`: 链路追踪模块.
    * `upload`: 文件上传模块.
    * `util`: 其它工具模块.
* `storage`: 文件存储目录.
    * `logs`: 日志目录.
    * `uploads`: 上传的文件(图片)目录.

* 注: 以上的目录结构基本按书作者的源码进行组织, 个人仅做了少量的修改. 其中目录命名的理解也是基于原书以及个人理解, 可能有不到位之处. 对于目录这样组织的结构是否合理, 个人经验较少, 在此没有太多改动.   
其中也有不太理解之处, 比如 `pkg/setting`等与该项目的相关性也比较大, 而作者并未放入 `internal` 路径下; `pkg/convert` 的类型转换以及 `pkg/email`等模块实现也比较简单但并没有将其合并放于 `pkg/util` 下.

## 基本框架
本项目选用 [Gin](https://github.com/gin-gonic/gin) 框架进行开发. Gin 是由 Go 语言编写的 Web 框架.  
对于该框架在此不多赘述, 在本项目中, 对请求路由处理, 响应以及中间件的设置都是基于 Gin 框架的.  

## 编译运行
在根目录下使用如下命令编译：
```shell
$ go build mian.go
```
