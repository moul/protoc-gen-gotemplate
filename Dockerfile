FROM		znly/protoc

ENV		GOPATH=/go \
		PATH=/go/bin:${PATH}

# Install deps and common tools
RUN		apk --update add make git go rsync libc-dev \
 &&		go get -u golang.org/x/tools/cmd/goimports

# Install protoc-gen-gotemplate
COPY		.	/go/src/github.com/moul/protoc-gen-gotemplate
WORKDIR		/go/src/github.com/moul/protoc-gen-gotemplate
RUN		go install .
