FROM znly/protoc
RUN apk --update add make git go rsync
COPY . /go/src/github.com/moul/protoc-gen-gotemplate
WORKDIR /go/src/github.com/moul/protoc-gen-gotemplate
RUN go install .
