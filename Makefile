TOP_PKG                      = github.com/operable/circuit
PKG_DIRS                    := $(shell find . \! -name "." -type d | sort)
PKGS                        := $(TOP_PKG) $(subst ., $(TOP_PKG), $(PKG_DIRS))

.PHONY: all test exe

all: test

test:
	go test -cover $(PKGS)
