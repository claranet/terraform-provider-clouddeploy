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

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

vendor-status:
	@govendor status

clean:
	$(RM) ${BINARY}

.PHONY: install fmt test testacc vet vendor-status clean
