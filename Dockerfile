FROM		znly/protoc:0.3.0

ENV		GOPATH=/go \
		PATH=/go/bin:${PATH}

# Install deps and common tools
RUN		apk --update add make git go rsync libc-dev \
 &&		go get -u golang.org/x/tools/cmd/goimports

# Install protoc-gen-gotemplate
COPY		. /go/src/moul.io/protoc-gen-gotemplate
WORKDIR		/go/src/moul.io/protoc-gen-gotemplate
RUN		git remote set-url origin https://github.com/moul/protoc-gen-gotemplate
RUN		go install . ./cmd/web-editor
