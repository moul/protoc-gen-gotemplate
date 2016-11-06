package main

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/kr/fs"
	"github.com/moul/funcmap"
)

type GenericTemplateBasedEncoder struct {
	templateDir string
	service     *descriptor.ServiceDescriptorProto
	file        *descriptor.FileDescriptorProto
}

type Ast struct {
	Filename string
	Service  *descriptor.ServiceDescriptorProto
	File     *descriptor.FileDescriptorProto
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

func (e *GenericTemplateBasedEncoder) buildContent(templateFilename string) (string, error) {
	fullPath := filepath.Join(e.templateDir, templateFilename)

	tmpl, err := template.New(templateFilename).Funcs(funcmap.FuncMap).ParseFiles(fullPath)
	if err != nil {
		return "", err
	}

	ast := Ast{
		Filename: templateFilename,
		Service:  e.service,
		File:     e.file,
	}

	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, ast); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (e *GenericTemplateBasedEncoder) Files() []*plugin_go.CodeGeneratorResponse_File {
	files := []*plugin_go.CodeGeneratorResponse_File{}

	templates, err := e.templates()
	if err != nil {
		panic(err)
	}

	for _, templateFilename := range templates {
		filename := templateFilename[0 : len(templateFilename)-len(".tmpl")]

		content, err := e.buildContent(templateFilename)
		if err != nil {
			panic(err)
		}

		files = append(files, &plugin_go.CodeGeneratorResponse_File{
			Content: &content,
			Name:    &filename,
		})
	}

	return files
}
