#!/usr/bin/env bash
#vs code 使用需要先安装 go extension

#使用代理下载
export GOPROXY=https://goproxy.io
export GO111MODULE=on

#microsoft visual studio code推荐工具
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

#自动完成工具gocode
#需要先装第一个stamblerre，产生gocode并改名为gocode-gomod，
#stamblerre可以使用go的modules特性，代码fork自mdempsky
#然后装第二个mdempsky，产生gocode，
#stamblerre这个东西就是支持modules的gocode，单独使用不用改名，
#应该是vs code的问题，非要改名gocode-gomod他才识别
go get -u -v github.com/stamblerre/gocode
mv $GOPATH/bin/gocode  $GOPATH/bin/gocode-gomod
go get -u -v github.com/mdempsky/gocode

#go get -u -v github.com/alecthomas/gometalinter
#gometalinter --install
#这个不维护了，使用如下代替
go get -u -v github.com/golangci/golangci-lint

#debug工具
go get -u -v github.com/go-delve/delve/cmd/dlv