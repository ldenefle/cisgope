BIN = cisgope

.PHONY: all
all: clean build

.PHONY: build
build: deps
	go build -o build/$(BIN) .

.PHONY: deps
deps:
	command -v dep >/dev/null || go get -u github.com/golang/dep/cmd/dep
	dep ensure

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -rf build
	rm -rf vendor
	go clean

