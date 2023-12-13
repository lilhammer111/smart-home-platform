# cd cmd/rpc/user
cwgo server --type RPC --idl ../../../static/idl/rpc/user.thrift --module git.zqbjj.top/pet/services/cmd/rpc/user --service user_micro --template ../../../static/tpl/custom_rpc_server_tpl/

# cd cmd/rpc/device
