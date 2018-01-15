package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	ggdescriptor "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"

	pgghelpers "github.com/moul/protoc-gen-gotemplate/helpers"
)

var (
	registry *ggdescriptor.Registry // some helpers need access to registry
)

const (
	boolTrue  = "true"
	boolFalse = "false"
)

func main() {
	g := generator.New()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err = proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	g.CommandLineParameters(g.Request.GetParameter())

	// Parse parameters
	var (
		templateDir       = "./templates"
		destinationDir    = "."
		debug             = false
		templateType      = "service"
		singlePackageMode = false
	)
	if parameter := g.Request.GetParameter(); parameter != "" {
		for _, param := range strings.Split(parameter, ",") {
			parts := strings.Split(param, "=")
			if len(parts) != 2 {
				log.Printf("Err: invalid parameter: %q", param)
				continue
			}
			switch parts[0] {
			case "template_dir":
				templateDir = parts[1]
			case "destination_dir":
				destinationDir = parts[1]
			case "single-package-mode":
				switch strings.ToLower(parts[1]) {
				case boolTrue, "t":
					singlePackageMode = true
				case boolFalse, "f":
				default:
					log.Printf("Err: invalid value for single-package-mode: %q", parts[1])
				}
			case "debug":
				switch strings.ToLower(parts[1]) {
				case boolTrue, "t":
					debug = true
				case boolFalse, "f":
				default:
					log.Printf("Err: invalid value for debug: %q", parts[1])
				}
			case "all":
				switch strings.ToLower(parts[1]) {
				case boolTrue, "t":
					templateType = "file"
				case boolFalse, "f":
				default:
					log.Printf("Err: invalid value for all: %q", parts[1])
				}
			case "type":
				templateType = parts[1]
			default:
				log.Printf("Err: unknown parameter: %q", param)
			}
		}
	}

	tmplMap := make(map[string]*plugin_go.CodeGeneratorResponse_File)
	concatOrAppend := func(file *plugin_go.CodeGeneratorResponse_File) {
		if val, ok := tmplMap[file.GetName()]; ok {
			*val.Content += file.GetContent()
		} else {
			tmplMap[file.GetName()] = file
			g.Response.File = append(g.Response.File, file)
		}
	}

	appendOnce := func(file *plugin_go.CodeGeneratorResponse_File) {
		if _, ok := tmplMap[file.GetName()]; ok {
			return
		}
		tmplMap[file.GetName()] = file
		g.Response.File = append(g.Response.File, file)
	}

	if singlePackageMode {
		registry = ggdescriptor.NewRegistry()
		pgghelpers.SetRegistry(registry)
		if err = registry.Load(g.Request); err != nil {
			g.Error(err, "registry: failed to load the request")
		}
	}

	files := g.Request.GetProtoFile()

	// Generate the encoders
	for _, file := range files {
		if !inStringSlice(file.GetName(), g.Request.FileToGenerate) {
			continue
		}
		switch templateType {
		case "none":
			if singlePackageMode {
				if _, err = registry.LookupFile(file.GetName()); err != nil {
					g.Error(err, "registry: failed to lookup file %q", file.GetName())
				}
			}
			encoder := NewGenericAllTemplateBasedEncoder(templateDir, files, debug, destinationDir)
			for _, tmpl := range encoder.Files() {
				appendOnce(tmpl)
			}
		case "file":
			if singlePackageMode {
				if _, err = registry.LookupFile(file.GetName()); err != nil {
					g.Error(err, "registry: failed to lookup file %q", file.GetName())
				}
			}
			encoder := NewGenericTemplateBasedEncoder(templateDir, file, debug, destinationDir)
			for _, tmpl := range encoder.Files() {
				concatOrAppend(tmpl)
			}
		case "message":
			for _, message := range file.GetMessageType() {
				encoder := NewGenericMessageTemplateBasedEncoder(templateDir, message, file, debug, destinationDir)
				for _, tmpl := range encoder.Files() {
					concatOrAppend(tmpl)
				}
			}
		case "service":
			for _, service := range file.GetService() {
				encoder := NewGenericServiceTemplateBasedEncoder(templateDir, service, file, debug, destinationDir)
				for _, tmpl := range encoder.Files() {
					concatOrAppend(tmpl)
				}
			}
		default:
			log.Printf("Err: invalid value for type: %q", templateType)
		}
	}

	// Generate the protobufs
	g.GenerateAllFiles()

	data, err = proto.Marshal(g.Response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}

	_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}

func inStringSlice(needle string, haystack []string) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}
