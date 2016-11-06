package main

import (
	"path/filepath"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/kr/fs"
)

type GenericTemplateBasedEncoder struct {
	templateDir string
	service     *descriptor.ServiceDescriptorProto
	file        *descriptor.FileDescriptorProto
}

func NewGenericTemplateBasedEncoder(templateDir string, service *descriptor.ServiceDescriptorProto, file *descriptor.FileDescriptorProto) (e *GenericTemplateBasedEncoder) {
	e = &GenericTemplateBasedEncoder{
		service:     service,
		file:        file,
		templateDir: templateDir,
	}
	return
}

func (e *GenericTemplateBasedEncoder) templates() ([]string, error) {
	filenames := []string{}

	walker := fs.Walk(e.templateDir)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			return nil, err
		}

		if walker.Stat().IsDir() {
			continue
		}

		if filepath.Ext(walker.Path()) != ".tmpl" {
			continue
		}

		rel, err := filepath.Rel(e.templateDir, walker.Path())
		if err != nil {
			return nil, err
		}

		filenames = append(filenames, rel)
	}

	return filenames, nil
}

func (e *GenericTemplateBasedEncoder) Files() []*plugin_go.CodeGeneratorResponse_File {
	files := []*plugin_go.CodeGeneratorResponse_File{}

	templates, err := e.templates()
	if err != nil {
		panic(err)
	}

	for _, templateFilename := range templates {
		filename := templateFilename[0 : len(templateFilename)-len(".tmpl")]

		content := "hello world"
		files = append(files, &plugin_go.CodeGeneratorResponse_File{
			Content: &content,
			Name:    &filename,
		})
	}

	return files
}
