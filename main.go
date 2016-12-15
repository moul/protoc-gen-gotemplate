package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func main() {
	g := generator.New()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	g.CommandLineParameters(g.Request.GetParameter())

	// Parse parameters
	templateDir := "./templates"
	destinationDir := "."
	debug := false
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
				break
			case "destination_dir":
				destinationDir = parts[1]
				break
			case "debug":
				switch strings.ToLower(parts[1]) {
				case "true", "t":
					debug = true
				case "false", "f":
				default:
					log.Printf("Err: invalid value for debug: %q", parts[1])
				}
				break
			default:
				log.Printf("Err: unknown parameter: %q", param)
			}
		}
	}

	// Generate the encoders
	for _, file := range g.Request.GetProtoFile() {
		for _, service := range file.GetService() {
			encoder := NewGenericTemplateBasedEncoder(templateDir, service, file, debug, destinationDir)
			g.Response.File = append(g.Response.File, encoder.Files()...)
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
