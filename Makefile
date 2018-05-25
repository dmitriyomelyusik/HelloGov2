lint:
	gometalinter --disable gas .
build: 
	go build -o bin/hello.run