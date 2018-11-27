#! /bin/bash
curl -X POST http://localhost:8765/file/upload -F "file=@//Users/pengbaojiang/pengbaojiang/code/gosrc/futty_golang/text.txt" -F "container=data" -F "key=hello/new.txt" -H "Content-Type: multipart/form-data"
