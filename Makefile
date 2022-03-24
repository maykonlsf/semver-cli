TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: test

lint:
	@if [ ! -f bin/golangci-lint ]; then \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.21.0; \
	fi
	./bin/golangci-lint run ./...

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/fmtcheck.sh'"

cover:
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	go test $(TEST) -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

test: fmtcheck
	@sh -c "go test ./... -timeout=2m -parallel=4"

build:
	@go build -o ./semver ./cmd/semver

build-linux-amd64:
	@mkdir -p ./dist/linux-amd64
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/linux-amd64/semver ./cmd/semver
	@echo "./dist/linux-amd64/semver (linux/amd64)"

build-windows-amd64:
	@mkdir -p ./dist/windows-amd64
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/windows-amd64/semver.exe ./cmd/semver
	@echo "./dist/windows-amd64/semver.exe (windows/amd64)"

build-macos-arm64:
	@mkdir -p ./dist/macos-arm64
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ./dist/macos-arm64/semver ./cmd/semver
	@echo "./dist/darwin-arm64/semver (darwin/arm64 - MacOS M1)"

build-macos-amd64:
	@mkdir -p ./dist/macos-amd64
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./dist/macos-amd64/semver ./cmd/semver
	@echo "./dist/darwin-amd64/semver (darwin/amd64 - MacOS Intel)"

build-all: build build-linux-amd64 build-windows-amd64 build-macos-arm64 build-macos-amd64

.PHONY: default test cover fmt fmtcheck lint build
