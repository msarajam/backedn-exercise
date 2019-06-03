ENTRY_POINT ?= ./cmd/server/main.go
PKGS ?= cmd pkg
EXPANDED_PKGS ?= ./cmd/... ./pkg/...

run: vendor
	@go run $(ENTRY_POINT)

test: vendor
	@go test $(EXPANDED_PKGS)

format:
	@gofmt -l -s -w $(PKGS)

clean:
	@rm -rf vendor

vendor:
	@dep ensure

.PHONY: run test format clean vendor
