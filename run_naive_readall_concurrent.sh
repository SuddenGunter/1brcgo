#!/bin/bash

cp ./naive_readall_concurrent/main.go ./main.go
go build
time GOGC=off ./1brcgo