#!/usr/bin/env bash
# 通过使用 file协议，可以简单的加载先前已经在本地缓冲好的依赖

# 注意：windows 下使用file协议运行该脚本时，最好在GOPATH同一个盘符下运行，否则可能会提示
# go get github.com/go-delve/delve/cmd/dlv: 
# open /go/pkg/mod/cache/download/github.com/go-delve/delve/cmd/dlv/@v/list: The system cannot find the path specified.
# 可以看见访问的路径中 盘符信息丢失了，具体原因不明。

# windwos bug?, file协议本来应该是 file:// / (3个=协议自身俩个+一个表示根目录的)，
# 这里却使用俩个就可以了，而且使用三个会报错，
# go get github.com/go-delve/delve/cmd/dlv: 
# open /f:/tmp/offline/mod/cache/download/github.com/go-delve/delve/cmd/dlv/@v/list: The filename, directory name, or volume label syntax is incorrect.
# 可见win上 并不需要表示根的/，直接使用盘符表示根即可， 但是操作系统和浏览器却可以打开 ///(3个的)，比较奇怪 
# file:///f:/go/pkg/mod/cache/download/github.com/go-delve/delve/@v/
export GOPROXY=file://$GOPATH/pkg/mod/cache/download

# 这个可以不在module化的 go 项目设置，有如下表现，
# module化的 go 项目中可以定位到本地依赖，而非模块化的 go项目中会去github服务器进行依赖拉取
# 原因可能是 module化的 go project 在其环境变量中存在 set GOMOD=F:\JetBrains\GoLand\goproxy\go.mod
# 这个是否会隐式的设置了 GO111MODULE=on
export GO111MODULE=on

go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v github.com/mdempsky/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v golang.org/x/tools/cmd/godoc
go get -u -v github.com/zmb3/gogetdoc
go get -u -v golang.org/x/lint/golint
go get -u -v github.com/fatih/gomodifytags
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/goimports
go get -u -v github.com/cweill/gotests/...
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v github.com/josharian/impl
go get -u -v github.com/haya14busa/goplay/cmd/goplay
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct

go get -u -v github.com/stamblerre/gocode
mv $GOPATH/bin/gocode $GOPATH/bin/gocode-gomod
go get -u -v github.com/mdempsky/gocode

go get -u -v github.com/golangci/golangci-lint

go get -u -v github.com/go-delve/delve/cmd/dlv