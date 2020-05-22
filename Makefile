APP_VERSION:=${shell date +%s}

default: build docker run

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-poc-linux-amd64 -ldflags="-s -w -X main.version=${APP_VERSION}" .

docker:
	docker build -t apipoc:${APP_VERSION} .

run:
	docker run -p 8443:8443 -it apipoc:${APP_VERSION}