export GO111MODULE=on
NAME=adbkitting
ENTRYPOINT=./

.PHONY: setup
setup:
	go mod download

.PHONY: build
build: setup
	GOOS=darwin GOARCH=amd64 go build -o dist/${NAME} ${ENTRYPOINT}
	GOOS=windows GOARCH=amd64 go build -o dist/${NAME}.exe ${ENTRYPOINT}
