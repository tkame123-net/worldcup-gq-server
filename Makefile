.PHONY: generate-wire gen test build run-local

generate-wire:
	cd di && wire

gen: generate-wire

test: gen
	echo "test tobe"
#	cp .env.test .env
#	ENV_FILE=/Volumes/SSD/rep/blog-server/.env go test ./...

build: test
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o build/user_api server.go

build-local: test
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o build/user_api server.go

run-local: build-local
	cp .env.dev .env
	build/user_api

run-linux: build
	cp .env.dev .env
	build/user_api