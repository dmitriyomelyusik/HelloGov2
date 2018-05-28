test:
	go test
lint:
	gometalinter --disable gas .
build:
	go build -o bin/hello helloworld.go
