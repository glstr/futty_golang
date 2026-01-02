#! /bin/bash

go build -o bin/snow main.go
rm -f nohup
killall  snow
nohup bin/snow &
