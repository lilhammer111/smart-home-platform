#!/bin/bash

# specify the service you need to update and the idl file
service_name="pet_api"
file="user.thrift"


# get the directory where the script file is located and change dir to the abs path
script_abs_path=$(realpath "$0")
echo "脚本绝对路径： ${script_abs_path} @@@"
script_dir=$(dirname "$script_abs_path")
cd "$script_dir" || exit

# change dir to http service dir
cd ../../cmd/http || exit

# update the http service
cwgo server \
    --type HTTP --service "$service_name" \
    --module "git.zqbjj.top/pet/services/cmd/http" \
    --idl "../../static/idl/http/$file" \
    --template "../../static/tpl/custom_http_server_tpl" \
    --pass "--handler_dir ./api/" \
    --pass "--router_dir ./router/" \
    --pass "--model_dir ./dto/hertz_gen/" \
    --pass "--snake_tag"

go mod tidy