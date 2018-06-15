#!/bin/bash

cd ..
sh ./autoUpdate.sh

cd admin
go build main.go
./main