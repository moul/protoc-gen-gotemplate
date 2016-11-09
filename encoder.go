package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/kr/fs"
	"github.com/moul/funcmap"
)

type GenericTemplateBasedEncoder struct {
	templateDir string
	service     *descriptor.ServiceDescriptorProto
	file        *descriptor.FileDescriptorProto
	debug       bool
}

type Ast struct {
	BuildDate     time.Time                          `json:"build-date"`
	BuildHostname string                             `json:"build-hostname"`
	BuildUser     string                             `json:"build-user"`
	GoPWD         string                             `json:"go-pwd",omitempty`
	PWD           string                             `json:"pwd"`
	Debug         bool                               `json:"debug"`
	File          *descriptor.FileDescriptorProto    `json:"file"`
	RawFilename   string                             `json:"raw-filename"`
	Filename      string                             `json:"filename"`
	Service       *descriptor.ServiceDescriptorProto `json:"service"`
}

func NewGenericTemplateBasedEncoder(templateDir string, service *descriptor.ServiceDescriptorProto, file *descriptor.FileDescriptorProto, debug bool) (e *GenericTemplateBasedEncoder) {
	e = &GenericTemplateBasedEncoder{
		service:     service,
		file:        file,
		templateDir: templateDir,
		debug:       debug,
	}

	if debug {
		log.Printf("new encoder: file=%q service=%q template-dir=%q", file.GetName(), service.GetName(), templateDir)
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
		if e.debug {
			log.Printf("new template: %q", rel)
		}

		filenames = append(filenames, rel)
	}

	return filenames, nil
}

func (e *GenericTemplateBasedEncoder) genAst(templateFilename string) (*Ast, error) {
	// prepare the ast passed to the template engine
	hostname, _ := os.Hostname()
	pwd, _ := os.Getwd()
	goPwd := ""
	if os.Getenv("GOPATH") != "" {
		goPwd, _ = filepath.Rel(os.Getenv("GOPATH")+"/src", pwd)
		if strings.Contains(goPwd, "../") {
			goPwd = ""
		}
	}
	ast := Ast{
		BuildDate:     time.Now(),
		BuildHostname: hostname,
		BuildUser:     os.Getenv("USER"),
		PWD:           pwd,
		GoPWD:         goPwd,
		File:          e.file,
		RawFilename:   templateFilename,
		// Filename:      "",
		Service: e.service,
	}
	return &ast, nil
}

func (e *GenericTemplateBasedEncoder) buildContent(templateFilename string) (string, error) {
	// initialize template engine
	fullPath := filepath.Join(e.templateDir, templateFilename)
	templateName := filepath.Base(fullPath)
	tmpl, err := template.New(templateName).Funcs(funcmap.FuncMap).ParseFiles(fullPath)
	if err != nil {
		return "", err
	}

	ast, err := e.genAst(templateFilename)
	if err != nil {
		return "", err
	}

	// translate the filename
	ast.Filename, err = e.translateString(templateFilename, templateFilename)
	if err != nil {
		return "", err
	}

	// generate the content
	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, ast); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (e *GenericTemplateBasedEncoder) translateString(input string, templateFilename string) (string, error) {
	buffer := new(bytes.Buffer)
	tmpl, err := template.New("").Funcs(funcmap.FuncMap).Parse(input)
	if err != nil {
		return "", err
	}

	ast, err := e.genAst(templateFilename)
	if err != nil {
		return "", err
	}

	if err := tmpl.Execute(buffer, ast); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (e *GenericTemplateBasedEncoder) Files() []*plugin_go.CodeGeneratorResponse_File {
	files := []*plugin_go.CodeGeneratorResponse_File{}

	templates, err := e.templates()
	if err != nil {
		log.Fatalf("cannot get templates from %q: %v", e.templateDir, err)
	}

	for _, templateFilename := range templates {
		content, err := e.buildContent(templateFilename)
		if err != nil {
			panic(err)
		}

		translatedFilename, err := e.translateString(templateFilename, templateFilename)
		if err != nil {
			panic(err)
		}
		filename := translatedFilename[0 : len(translatedFilename)-len(".tmpl")]

		files = append(files, &plugin_go.CodeGeneratorResponse_File{
			Content: &content,
			Name:    &filename,
		})
	}

	return files
}
