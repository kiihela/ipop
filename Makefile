TAGS ?= "sqlite"
GO_BIN ?= go
SODA_BIN ?= ~/go/bin/soda

install: deps

deps:
	$(GO_BIN) get -tags ${TAGS} ./...
	$(GO_BIN) get -tags ${TAGS} github.com/gobuffalo/pop/soda
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif

db:
	$(SODA_BIN) help

test:
	$(GO_BIN) test -tags ${TAGS} ./...

ci-test:
	$(GO_BIN) test -tags ${TAGS} -race ./...

lint:
	gometalinter --vendor ./... --deadline=1m --skip=internal

update:
	$(GO_BIN) get -u -tags ${TAGS}
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif
	make test
	make install
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif
