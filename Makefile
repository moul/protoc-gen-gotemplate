.PHONY: build
build:
	go build -o protoc-gen-gotemplate .

.PHONY: install
install:
	go install .

.PHONY: test
test:	build
	cd examples/dummy && make
	cd examples/go-kit && make
