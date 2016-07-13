#
# Simple Makefile
#

build:
	go build -o bin/mergepath cmds/mergepath/mergepath.go 

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f mergepath-binary-release.zip ]; then rm -f mergepath-binary-release.zip; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/mergepath/mergepath.go

release:
	./mk-release.sh
