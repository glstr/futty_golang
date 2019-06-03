#! bin/bash
flag=$1
if [ -z $flag ];then
    echo "usage:"    
    echo "sh build_server.sh server"
    echo "sh build_server.sh rpcxserver"
    echo "sh build_server.sh tcpserver"
fi

if [[ $flag == 'server' ]];then
    go build -o bin/server src/main/server.go
elif [[ $flag == 'rpcxserver' ]];then
    go build -o bin/rpcxserver src/server/rpcxserver/rpcxserver.go
elif [[ $flag == 'tcpserver' ]];then
    go build -o bin/tcpserver src/server/tcpserver/main/main.go
else
    echo "invalid param"
fi
