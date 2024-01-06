APP=$(shell basename $(shell git remote get-url origin))
REGISTRY=ghcr.io
REPOSYTORY=vasyl-babynets
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux
TARGETARCH=amd64

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

get:
	go get

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags  "-X="github.dev/vasyl-babynets/kbot/cmd.appVersion=${VERSION}

linux: TARGETOS=linux 
linux: build image

windows: TARGETOS=windows
windows: build image

macOS: TARGETOS=darwin
macOS: TARGETARCH=arm64
macOS: build image

image:
	docker build . -t ${REGISTRY}/${REPOSYTORY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

push:
	docker push ${REGISTRY}/${REPOSYTORY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

push-linux: push

push-windows: push

push-macOS: TARGETARCH=arm64
push-macOS: push

clean:
	docker rmi -f ${REGISTRY}/${REPOSYTORY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

clean-linux: clean

clean-windows: clean

clean-macOS: TARGETARCH=arm64
clean-macOS: clean