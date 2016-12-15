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
)

type GenericTemplateBasedEncoder struct {
	templateDir    string
	service        *descriptor.ServiceDescriptorProto
	file           *descriptor.FileDescriptorProto
	debug          bool
	destinationDir string
}

type Ast struct {
	BuildDate      time.Time                          `json:"build-date"`
	BuildHostname  string                             `json:"build-hostname"`
	BuildUser      string                             `json:"build-user"`
	GoPWD          string                             `json:"go-pwd,omitempty"`
	PWD            string                             `json:"pwd"`
	Debug          bool                               `json:"debug"`
	DestinationDir string                             `json:"destination-dir"`
	File           *descriptor.FileDescriptorProto    `json:"file"`
	RawFilename    string                             `json:"raw-filename"`
	Filename       string                             `json:"filename"`
	TemplateDir    string                             `json:"template-dir"`
	Service        *descriptor.ServiceDescriptorProto `json:"service"`
	Environment    []string                           `json:"environment"`
}

func NewGenericTemplateBasedEncoder(templateDir string, service *descriptor.ServiceDescriptorProto, file *descriptor.FileDescriptorProto, debug bool, destinationDir string) (e *GenericTemplateBasedEncoder) {
	e = &GenericTemplateBasedEncoder{
		service:        service,
		file:           file,
		templateDir:    templateDir,
		debug:          debug,
		destinationDir: destinationDir,
	}

	if debug {
		log.Printf("new encoder: file=%q service=%q template-dir=%q", file.GetName(), service.GetName(), templateDir)
	}

	return
}

func (e *GenericTemplateBasedEncoder) templates() ([]string, error) {
	filenames := []string{}

	err := filepath.Walk(e.templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".tmpl" {
			return nil
		}
		rel, err := filepath.Rel(e.templateDir, path)
		if err != nil {
			return err
		}
		if e.debug {
			log.Printf("new template: %q", rel)
		}
		filenames = append(filenames, rel)
		return nil
	})
	return filenames, err
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
		BuildDate:      time.Now(),
		BuildHostname:  hostname,
		BuildUser:      os.Getenv("USER"),
		PWD:            pwd,
		GoPWD:          goPwd,
		File:           e.file,
		TemplateDir:    e.templateDir,
		DestinationDir: e.destinationDir,
		RawFilename:    templateFilename,
		Filename:       "",
		Environment:    os.Environ(),
		Service:        e.service,
	}
	buffer := new(bytes.Buffer)
	tmpl, err := template.New("").Funcs(ProtoHelpersFuncMap).Parse(templateFilename)
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buffer, ast); err != nil {
		return nil, err
	}
	ast.Filename = buffer.String()
	return &ast, nil
}

func (e *GenericTemplateBasedEncoder) buildContent(templateFilename string) (string, string, error) {
	// initialize template engine
	fullPath := filepath.Join(e.templateDir, templateFilename)
	templateName := filepath.Base(fullPath)
	tmpl, err := template.New(templateName).Funcs(ProtoHelpersFuncMap).ParseFiles(fullPath)
	if err != nil {
		return "", "", err
	}

	ast, err := e.genAst(templateFilename)
	if err != nil {
		return "", "", err
	}

	// generate the content
	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, ast); err != nil {
		return "", "", err
	}

	return buffer.String(), ast.Filename, nil
}

func (e *GenericTemplateBasedEncoder) Files() []*plugin_go.CodeGeneratorResponse_File {
	templates, err := e.templates()
	if err != nil {
		log.Fatalf("cannot get templates from %q: %v", e.templateDir, err)
	}

	length := len(templates)
	files := make([]*plugin_go.CodeGeneratorResponse_File, 0, length)
	errChan := make(chan error, length)
	resultChan := make(chan *plugin_go.CodeGeneratorResponse_File, length)
	for _, templateFilename := range templates {
		go func(tmpl string) {
			content, translatedFilename, err := e.buildContent(tmpl)
			if err != nil {
				errChan <- err
				return
			}
			filename := translatedFilename[:len(translatedFilename)-len(".tmpl")]

			resultChan <- &plugin_go.CodeGeneratorResponse_File{
				Content: &content,
				Name:    &filename,
			}
		}(templateFilename)
	}
	for i := 0; i < length; i++ {
		select {
		case f := <-resultChan:
			files = append(files, f)
		case err = <-errChan:
		}
	}
	if err != nil {
		panic(err)
	}
	return files
}
