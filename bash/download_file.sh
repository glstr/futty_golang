#! /bin/bash

curl -X POST http://localhost:8765/file/download -d '{"key":"hello/new.txt","container":"data"}'
