#
# Simple Makefile
#

build: mergepath.go
	go build -o bin/mergepath mergepath.go 

clean: mergepath
	rm bin/mergepath

install: mergepath.go
	go install mergepath.go

