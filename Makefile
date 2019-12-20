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
	CGO_ENABLED=0
	GOOS=linux
	@go build -o semver ./cmd/semver

.PHONY: default test cover fmt fmtcheck lint
