.PHONY: re
re: clean build test

.PHONY: build
build:
	@mkdir -p output

	@# proto-gen-go
	protoc -I./proto --go_out=plugins=grpc:output proto/aaa/aaa.proto
	protoc -I./proto --go_out=plugins=grpc:output proto/bbb/bbb.proto
	@rm -rf output/aaa output/bbb
	@mv output/github.com/moul/protoc-gen-gotemplate/examples/single-package-mode/output/* output/
	@rm -rf output/github.com

	@# protoc-gen-gotemplate
	protoc -I./proto --gotemplate_out=template_dir=templates,single-package-mode=true:output proto/bbb/bbb.proto
	gofmt -w .

.PHONY: test
test:
	go test -i ./output/...
	go test -v ./output/...

.PHONY: clean
clean:
	rm -rf output
