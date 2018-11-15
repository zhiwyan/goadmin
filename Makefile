#!/bin/bash
# jun.guo@wenba100.com

CURDIR:=$(shell pwd)
GO ?= go

all:
	git pull
	$(GO) build -o bin/config_server main.go
#debug: gdb check

init :
	echo init
	#git clone https://gitlab-wenba.xueba100.com:2443/fudao2/vendor.git
	# go get github.com/gin-gonic/gin
	# go get github.com/go-ini/ini
	# go get github.com/jinzhu/gorm
	# go get github.com/jinzhu/gorm/dialects/mysql
	# go get github.com/gomodule/redigo/redis
	mkdir ./logs
	chmod -R 0777 ./logs

clean:
	rm -rf bin/classroom

.PHONY : clean
