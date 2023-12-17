#!/bin/bash

# specify the service you need to update and the idl file
service_name="micro_user"
file="user.thrift"
dir="user"

# get the directory where the script file is located and change dir to the abs path
script_abs_path=$(realpath "$0")
echo "脚本绝对路径： ${script_abs_path} @@@"
script_dir=$(dirname "$script_abs_path")
cd "$script_dir" || exit

# change dir to http service dir
cd ../../cmd/rpc/$dir || exit

# update
cwgo server --type RPC \
            --idl ../../../static/idl/rpc/$file \
            --module git.zqbjj.top/pet/services/cmd/rpc/$dir \
            --service $service_name \
            --template ../../../static/tpl/custom_rpc_server_tpl/ \
            --pass "--thrift code_ref"

go mod tidy