PROJECT = github.com/albertrdixon/amnd
REV ?= $$(git rev-parse --short=8 HEAD)
BRANCH ?= $$(git rev-parse --abbrev-ref HEAD | tr / _)
EXECUTABLE = amnd
BINARY = cmd/amnd/main.go
LDFLAGS = "-s -X $(PROJECT)/amnd.SHA $(REV)"
TEST_COMMAND = godep go test
PLATFORMS = linux darwin

.PHONY: dep-save dep-restore test test-verbose build install

all: dep-restore test build package

help:
	@echo "Available targets:"
	@echo ""
	@echo "  dep-save       : Save dependencies (godep save)"
	@echo "  dep-restore    : Restore dependencies (godep restore)"
	@echo "  test           : Run package tests"
	@echo "  test-verbose   : Run package tests with verbose output"
	@echo "  build          : Build binary (go build)"
	@echo "  package        : Create binary tar"
	@echo "  clean          : Clean working directory"

dep-save:
	@echo "==> Saving dependencies to ./Godeps"
	@godep save ./...

dep-restore:
	@echo "==> Restoring dependencies from ./Godeps"
	@godep restore

test:
	@echo "==> Running all tests"
	@echo ""
	@$(TEST_COMMAND) ./...

test-verbose:
	@echo "==> Running all tests (verbose output)"
	@ echo ""
	@$(TEST_COMMAND) -test.v ./...

build: clean
	@echo "==> Building $(EXECUTABLE) with ldflags '$(LDFLAGS)'"
	@ GOOS=linux CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags $(LDFLAGS) -o bin/$(EXECUTABLE)-linux $(BINARY)
	@ GOOS=darwin CGO_ENABLED=0 godep go build -a -ldflags $(LDFLAGS) -o bin/$(EXECUTABLE)-darwin $(BINARY)

package: build
	@echo "==> Tar'ing up the binaries"
	@for p in $(PLATFORMS) ; do \
		echo "==> Tar'ing up $$p/amd64 binary" ; \
		test -f bin/$(EXECUTABLE)-$$p && \
		cp -f bin/$(EXECUTABLE)-$$p $(EXECUTABLE) && \
		tar czf $(EXECUTABLE)-$$p.tar.gz $(EXECUTABLE) && \
		rm -f $(EXECUTABLE) ; \
	done

clean:
	@echo "==> Cleaning working directory"
	@go clean ./...
	rm -vf $(EXECUTABLE) *.tar.gz bin/*
