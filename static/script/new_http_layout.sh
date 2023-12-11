#!/bin/bash

# specify the project core variables
GO_MODULE="git.zqbjj.top/pet/services/pet-feeder/cmd/http"

# get the directory where the script file is located
script_abs_path=$(realpath "$0")
echo "脚本绝对路径： ${script_abs_path} @@@"
script_dir=$(dirname "$script_abs_path")

# cd project root dir and init go workspace
cd "$script_dir" || exit
cd ../../
go work init
#ROOT_DIR=$(pwd)

# mkdir of http
HTTP_RELA_DIR="cmd/http"

mkdir -p $HTTP_RELA_DIR


# define variables of output directories and static directory
HTTP_OUT_DIR=$(cd cmd/http||exit ;pwd)
echo "http绝对目录是 $HTTP_OUT_DIR"
STATIC_DIR=$(cd static || exit ; pwd)
#echo "static绝对目录是$STATIC_DIR"

# loop the following directory for generating the layout including each http server
cd $HTTP_RELA_DIR || exit
go mod init $GO_MODULE


# make sure that cwgo has been installed
if ! command -v cwgo &> /dev/null
then
    echo "cwgo 未安装，正在安装..."
    go install github.com/cloudwego/cwgo@latest
else
    echo "cwgo 已安装"
fi

for file in "$STATIC_DIR"/idl/http/*
do
  # make sure that the $file is a real file
  if [ -f "$file" ]
  then
    # remove the file path and suffix extension .thrift
    service_name=$(basename "$file" .thrift)
    echo "服务名： $service_name"
    # new server layout by cwgo
    if [ "$service_name" != "std_resp" ]
    then
    cwgo server \
        --type HTTP --service "$service_name" \
        --module "$GO_MODULE" \
        --idl "$file" \
        --template "$STATIC_DIR/tpl/custom_http_server_tpl" \
        --pass "--handler_dir ./api/" \
        --pass "--router_dir ./router/" \
        --pass "--model_dir ./dto/hertz_gen/"
    fi
  fi
done

go mod tidy

go work use .
