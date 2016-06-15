TOP_PKG                      = github.com/operable/circuit
BUILD_DIR                    = _build
PKG_DIRS                    := $(shell find . -not -path '*/\.*' -type d | grep -v ${BUILD_DIR} | sort)
PKGS                        := $(TOP_PKG) $(subst ., $(TOP_PKG), $(PKG_DIRS))

.PHONY: all test exe

all: test

test:
	go test -cover $(PKGS)
