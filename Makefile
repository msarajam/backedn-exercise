ENTRY_POINT ?= ./cmd/server/main.go
PKGS ?= cmd pkg
EXPANDED_PKGS ?= ./cmd/... ./pkg/...

run: vendor
	@go run $(ENTRY_POINT)

test: vendor
	@go test $(EXPANDED_PKGS)

format:
	@gofmt -l -s -w $(PKGS)

lint:
	@golangci-lint run

clean:
	@rm -rf vendor

vendor:
	@export Go111MODULE=On
	@go mod init
	@go mod tidy
	@go mod vendor

generate: vendor
	@go generate $(EXPANDED_PKGS)

bundle:
	@git bundle create backend-exercise_completed.bundle HEAD master

.PHONY: run test format lint clean vendor generate bundle
