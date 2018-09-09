### portal

## 目录结构

``` bash
.
├── config                    # 主要配置(mysql, redis)
│   ├── config.example.json
│   ├── config.go
│   └── config.json
├── controller                # 控制层
│   ├── app
│   │   └── app.go
│   ├── captcha
│   │   └── captcha.go
│   ├── home.go
│   ├── inter
│   │   └── inter.go
│   ├── menu
│   │   └── menu.go
│   ├── openAuth
│   │   └── openAuth.go
│   ├── permission
│   │   └── permission.go
│   ├── resource
│   │   └── resource.go
│   ├── role
│   │   └── role.go
│   └── user
│       └── user.go
├── database                  # model实体的sql
│   ├── app.go
│   ├── common.go
│   ├── db.go
│   ├── interface.go
│   ├── log.go
│   ├── openAuth.go
│   ├── resource.go
│   ├── role.go
│   ├── roleResource.go
│   ├── router.go
│   └── user.go
├── doc                       # 接口文档
│   └── api.md
├── error.log
├── Gopkg.lock
├── Gopkg.toml
├── main.go                   # main package
├── middleware                # 中间件
│   ├── auth.go
│   └── cors.go
├── model                     # model定义
│   ├── app.go
│   ├── common.go
│   ├── interface.go
│   ├── log.go
│   ├── openAuth.go
│   ├── resource.go
│   ├── role.go
│   ├── roleResource.go
│   ├── router.go
│   ├── token.go
│   └── user.go
├── README.md
├── router                    # 路由
│   └── router.go
├── service                   # 业务逻辑
│   ├── app.go
│   ├── captcha.go
│   ├── interface.go
│   ├── openAuth.go
│   ├── resource.go
│   ├── role.go
│   ├── roleResource.go
│   ├── router.go
│   └── user.go
├── test                      # 单元测试
│   ├── aes_test.go
│   ├── app_test.go
│   ├── const_test.go
│   ├── resource_test.go
│   ├── role_test.go
│   ├── route_test.go
│   ├── strconv_test.go
│   ├── token_test.go
│   ├── user_test.go
│   └── utils_test.go
└── util                      # 工具包
    ├── aes.go
    ├── const.go
    ├── http.go
    ├── reflect.go
    └── utils.go
```

## install dependence

``` bash
$ git clone git@github.com:yourlei/portal.git
$ cd portal 
$ dep ensure
```

## run

``` bash
$ go run main.go
```

监听3000端口, 测试服务启动成功:

``` bash
➜  ~ curl localhost:3000
{"code":0,"msg":"service running"}
```