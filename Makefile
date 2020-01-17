build:
	go build -o bin/client ./src/client.go
	go build -o bin/server ./src/server.go

clean:
	rm bin/*