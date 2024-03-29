.PHONY: test gen-ts build run-local

test:
	cp .env.test .env
	ENV_FILE=/Volumes/SSD/rep/blog-server/.env go test ./...

gen-ts:
	npm run generate

gen-schema:
	gqlgen generate

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o build/user_api server.go

build-local:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o build/user_api server.go

run-local: build-local
	cp .env.dev .env
	build/user_api

run-linux: build
	cp .env.dev .env
	build/user_api

deploy-dev:
	git push heroku develop:master
