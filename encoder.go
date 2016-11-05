package main

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
)

type GenericTemplateBasedEncoder struct {
	service *descriptor.ServiceDescriptorProto
	file    *descriptor.FileDescriptorProto
}

func NewGenericTemplateBasedEncoder(service *descriptor.ServiceDescriptorProto, file *descriptor.FileDescriptorProto) (e *GenericTemplateBasedEncoder) {
	e = &GenericTemplateBasedEncoder{
		service: service,
		file:    file,
	}
	return
}

func (e *GenericTemplateBasedEncoder) Files() []*plugin_go.CodeGeneratorResponse_File {
	//log.Printf("file: %v\n", e.file)
	//log.Printf("service: %v\n", e.service)
	var content string = "hello world"
	var fileName string = "test.txt"
	return []*plugin_go.CodeGeneratorResponse_File{
		&plugin_go.CodeGeneratorResponse_File{
			Content: &content,
			Name:    &fileName,
		},
	}
}
