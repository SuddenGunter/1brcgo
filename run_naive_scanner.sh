#!/bin/bash

cp ./naive_scanner/main.go ./main.go
go build
time GOGC=off ./1brcgo