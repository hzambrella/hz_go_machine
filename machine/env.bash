#/!bin/bash

# 设定git库地址转换, 以便解决部分包的库被墙的问题
#git config --global url."git@git.ot24.net:".insteadOf "https://git.ot24.net"
#git config --global url."https://github.com/golang/".insteadOf "https://go.googlesource.com/"
#PJROOT:为了让程序找到配置文件的路径。如果不用，程序的相对路径会出现无法找到文件的错误。
#supervisor可以取配置
#GOLIBS：包的路径
#GOPATH：不用说了
PWDDIR=`pwd`
#改名为PJROOT
export ETCDIR=$PWDDIR/etc
#配置文件目录
echo $ETCDIR
mkdir -p $ETCDIR
#MVC架构
#view
mkdir -p src    
cd src 
mkdir -p public
mkdir -p public/js
mkdir -p public/html
mkdir -p public/css
mkdir -p public/images
#control
mkdir -p routes
#model
mkdir -p model
cd ..
#第三方库和包,.gitignore 要忽略掉
export LIBDIR="$(dirname "$PWDDIR")/golibs"
mkdir -p $LIBDIR
#自己的 库和包
export LIB="$(dirname "$PWDDIR")/lib"
mkdir -p $LIB
#gopath
export GOPATH=$LIBDIR:$LIB:$PWDDIR
echo $GOPATH
