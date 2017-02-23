#!/bin/bash

# Go Path
# CURDIR=`pwd`
# OLDGOPATH="$GOPATH"
# export GOPATH="$CURDIR/../../"

LogPrefix=">>>>"

# 更新Bindata
while getopts "ath" arg
do
    case $arg in
         h)
            echo "- [h] [a] [t] help"
            echo "-h [a] help"
            echo "-a update assets bindata"
            echo "-t update template bindata"
            exit
            ;;
        a)
            echo -e "$LogPrefix `date +"%H:%M:%S"` [\033[44;37m update \033[0m] assets bindata"
            go-bindata -ignore=\\.DS_Store -debug -ignore=assets.go -pkg="assets" -o assets/assets.go assets/...
            ;;
        t)
            echo -e "$LogPrefix `date +"%H:%M:%S"` [\033[44;37m update \033[0m] template bindata"
            go-bindata -ignore=\\.DS_Store -debug -ignore=template.go -pkg="template" -o template/template.go template/...
            ;;
        ?)  #当有不认识的选项的时候arg为?
            echo "unkonw argument"
            echo "-h [a] help"
            exit 1
            ;;
    esac
done

# 清空pkg
echo -e "$LogPrefix `date +"%H:%M:%S"` rmove pkg"
rm -rf pkg/*

echo -e "$LogPrefix `date +"%H:%M:%S"` [\033[42;37m run \033[0m] server"
go run server.go