#!/bin/bash

# specify the service you need to update and the idl file
service_name="user_micro"
file="user.thrift"


# get the directory where the script file is located and change dir to the abs path
script_abs_path=$(realpath "$0")
echo "脚本绝对路径： ${script_abs_path} @@@"
script_dir=$(dirname "$script_abs_path")
cd "$script_dir" || exit

# change dir to http service dir
cd ../../cmd/http || exit

# update the http service
cwgo client --type RPC \
            --idl "../../static/idl/rpc/$file" \
            --template ../../static/tpl/custom_rpc_client_tpl/ \
            --service $service_name \
            --module git.zqbjj.top/pet/services/cmd/http

go mod tidy