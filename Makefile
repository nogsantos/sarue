XC_OS="linux darwin windows"
XC_ARCH="amd64"
XC_PARALLEL="2"
BIN="./bin"
SRC=$(shell find . -name "*.go")

ifeq (, $(shell which gox))
$(warning "could not find gox in $(PATH), run: go get github.com/mitchellh/gox")
endif

ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in $(PATH), run: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.0")
endif

.PHONY: fmt all build

default: all

all: build

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v

build:
	$(info ******************** building  ********************)
	mkdir -p $(BIN)
	gox \
		-os=$(XC_OS) \
		-arch=$(XC_ARCH) \
		-parallel=$(XC_PARALLEL) \
		-output=$(BIN)/{{.Dir}}_{{.OS}}_{{.Arch}} \
		;

clean:
	rm -rf $(BIN)
