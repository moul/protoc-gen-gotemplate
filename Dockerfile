FROM golang:1.7.3
COPY . /go/src/github.com/moul/protoc-gen-gotemplate
WORKDIR /go/src/github.com/moul/protoc-gen-gotemplate
RUN go install .
ENTRYPOINT ["protoc-gen-gotemplate"]
