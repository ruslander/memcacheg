build:
	go install transport

	go build -o bin/client ./src/client/client.go
	go build -o bin/server ./src/server/server.go

deps:
	go get github.com/codahale/hdrhistogram
	go install github.com/codahale/hdrhistogram

test:
	go test -bench=. ./src/...

clean:
	rm -rf pkg/*.*
	rm bin/*
