# don't pass --gen-path dto/
cwgo client --type RPC --idl ../../static/idl/rpc/user.thrift  --template ../../static/tpl/custom_rpc_client_tpl/ --service user_micro --module git.zqbjj.top/pet/services/cmd/http
