export GOPATH = $(shell pwd)


build:
	go install transport
	go build -o bin/client ./src/client/client.go
	go build -o bin/server ./src/server/server.go

test:
	go test -bench=. ./src/...

clean:
	rm -rf pkg/*.*
	rm bin/*
