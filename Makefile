# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)

# Setup name variables for the package/tool
AUTHOR := jlaswell
NAME := conerror
PKG := github.com/$(AUTHOR)/$(NAME)

.PHONY: docs
docs:
	@godoc2md -play $(PKG)  > README.md
.PHONY: lint
lint:
	@golint ./...

.PHONY: test
test:
	@go test -v -race ./...

.PHONY: vet
vet:
	@go vet ./...
