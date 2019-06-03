#! bin/bash
flag=$1
if [ -z $flag ];then
    echo "usage:"    
    echo "sh build_client.sh client"
    echo "sh build_client.sh rpcxclient"
    echo "sh build_client.sh tcpclient"
fi

if [[ $flag == 'client' ]];then
    go build -o bin/client src/client/client.go
elif [[ $flag == 'rpcxclient' ]];then
    go build -o bin/rpcxclient src/client/rpcxclient/main/main.go
elif [[ $flag == 'tcpclient' ]];then 
    go build -o bin/tcpclient src/client/tcpclient/main/main.go
else
    echo "invalid param"
fi

