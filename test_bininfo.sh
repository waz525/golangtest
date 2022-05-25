#!/usr/bin/env bash

#set -x

# 将 log 原始字符串中的单引号替换成双引号
Version="1.0.1"
# 检查源码在git commit 基础上，是否有本地修改，且未提交的内容
Author="Worden"
# 获取当前时间
BuildTime=`date +'%Y.%m.%d %H:%M:%S'`
# 获取 Go 的版本
BuildGoVersion=`go version`

# 将以上变量序列化至 LDFlags 变量中
LDFlags=" \
    -X 'BinInfo.Version=${Version}' \
    -X 'BinInfo.Author=${Author}' \
    -X 'BinInfo.BuildTime=${BuildTime}' \
    -X 'BinInfo.BuildGoVersion=${BuildGoVersion}' \
"

#echo "$LDFlags"

ROOT_DIR=`pwd`

go build -ldflags "$LDFlags" test_bininfo.go

./test_bininfo -v


