GOOS	 ?= linux
GOARCH  = amd64

export GO111MODULE ?= on

test:
	go test -v ./...

mod:
	go mod tidy

run:
	go run app/main.go

compile:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o encryptor app/main.go