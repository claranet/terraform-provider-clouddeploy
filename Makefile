SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go' | grep -v vendor)

BINARY=bin/terraform-provider-ghost

VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	go build ${LDFLAGS} -o ${BINARY} main.go

install: fmt
	go install

fmt:
	gofmt $(SOURCES)

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: install fmt clean
