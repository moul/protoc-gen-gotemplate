.PHONY: build
build:
	go build -v -i -o protoc-gen-gotemplate .

.PHONY: install
install:
	go install .
	go install ./cmd/web-editor

.PHONY: test
test:	install
	cd examples/time && make
	cd examples/enum && make
	cd examples/import && make
	cd examples/dummy && make
	cd examples/flow && make
	cd examples/concat && make
	cd examples/flow && make
	cd examples/sitemap && make
	cd examples/go-generate && make
	cd examples/single-package-mode && make
	cd examples/helpers && make
#	cd examples/go-kit && make

.PHONY: docker.build
docker.build:
	docker build --pull -t moul/protoc-gen-gotemplate .

.PHONY: docker.push
docker.push: docker.build
	docker push moul/protoc-gen-gotemplate

.PHONY: lint
lint:
	gometalinter --disable-all --enable=errcheck --enable=vet --enable=vetshadow --enable=golint --enable=gas --enable=ineffassign --enable=goconst --enable=goimports --enable=gofmt --exclude="Binds to all network interfaces" --exclude="should have comment" --enable=staticcheck --enable=gosimple --enable=misspell --deadline=120s . ./cmd/... ./helpers/...
