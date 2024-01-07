#!/bin/bash

cp ./naive_readall/main.go ./main.go
go build
time GOGC=off ./1brcgo