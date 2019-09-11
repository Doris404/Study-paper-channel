## Notes on pd 1 setup

#### GOPATH

GOPATH是go语言中重要的环境变量，我们可以通过在命令
行中输入```go env```查看go语言的环境

```
C:\>go env
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\用户名\AppData\Local\go-build
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOOS=windows
set GOPATH=C:\Users\用户名\go
set GOPROXY=
set GORACE=
set GOROOT=c:\go
set GOTMPDIR=
set GOTOOLDIR=c:\go\pkg\tool\windows_amd64
set GCCGO=gccgo
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
```
有两个常用的命令```go get```和```go install```

**go get**

* 从远程下载需要用到的包
* 执行go install

**go install**
* 生成可执行文件直接放在bin目录下
