SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go' | grep -v vendor)
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
	gofmt $(SOURCES)

clean:
	$(RM) ${BINARY}

.PHONY: install fmt clean
