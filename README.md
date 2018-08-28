### portal

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