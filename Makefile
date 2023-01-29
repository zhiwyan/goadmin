#!/bin/bash
# jun.guo@wenba100.com

CURDIR:=$(shell pwd)
GO ?= go

all:
	$(GO) build -o bin/goadmin main.go
#debug: gdb check

init :
	echo init
	mkdir ./logs
	chmod -R 0777 ./logs

clean:
	rm -rf bin/goadmin

.PHONY : clean
