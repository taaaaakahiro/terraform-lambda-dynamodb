run:
	go run ./pkg/cmd/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o ./build/bin/hello ./pkg/cmd/main.go
	zip ./build/zip/hello.zip ./build/bin/hello

.PHONY: build