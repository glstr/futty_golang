#! /bin/bash

echo $GOPATH
go get -u github.com/jstemmer/gotags
go get -u github.com/nsf/gocode
go get -u github.com/derekparker/delve/cmd/dlv
go get golang.org/x/tools/cmd/guru
go get golang.org/x/tools/cmd/goimports


