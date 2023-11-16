.PHONY: build
build:
	go mod tidy && CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-s -w" -o application .