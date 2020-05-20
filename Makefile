GOPKG ?= moul.io/protoc-gen-gotemplate
DOCKER_IMAGE ?= moul/protoc-gen-gotemplate
GOBINS ?= . ./cmd/web-editor
GOLIBS ?= .

all: test install

include rules.mk

.PHONY: examples
examples:	install
	cd examples/time && make
	cd examples/enum && make
	cd examples/import && make
	cd examples/dummy && make
	cd examples/flow && make
	cd examples/concat && make
	cd examples/flow && make
	cd examples/sitemap && make
	cd examples/go-generate && make
  #cd examples/single-package-mode && make
	cd examples/helpers && make
	cd examples/arithmetics && make
  #cd examples/go-kit && make
