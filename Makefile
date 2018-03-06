SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go' | grep -v vendor)
TEST?=./...
BIN_FOLDER=bin/
BINARY=$(BIN_FOLDER)terraform-provider-ghost
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

default: $(BINARY)

$(BINARY): $(SOURCES)
	go build ${LDFLAGS} -o ${BINARY} main.go

install: fmt
	go install

fmt:
	gofmt -w $(SOURCES)

test:
	go test $(TEST) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

clean:
	$(RM) ${BINARY}

.PHONY: install fmt test clean
