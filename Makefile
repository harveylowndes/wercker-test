PKG := github.com/harveylowndes/wercker-test
BUILD_DIR := dist
BIN_DIR := ${BUILD_DIR}/bin
BIN := bin

GOOS ?= darwin
GOARCH ?= amd64

BUILD := 1.0.0
VERSION ?= ${BUILD}

SRC_DIRS := cmd pkg 

.PHONY: clean
clean:
	rm -rf ${BUILD_DIR}

.PHONY: build
build:
	mkdir -p ${BIN_DIR}
	GOOS=${GOOS} \
	    CGO_ENABLED=0 \
	    GOARCH=${GOARCH} \
	    go build \
	    -i \
	    -v \
	    -ldflags="-s -w -X main.version=${VERSION} -X main.build=${BUILD}" \
	    -o ${BIN_DIR}/${BIN} ./cmd/oci

.PHONY: test
test:
	@./hack/test.sh $(SRC_DIRS)
	