package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/huandu/xstrings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	options "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
)

var ProtoHelpersFuncMap = template.FuncMap{
	"string": func(i interface {
		String() string
	}) string {
		return i.String()
	},
	"json": func(v interface{}) string {
		a, _ := json.Marshal(v)
		return string(a)
	},
	"prettyjson": func(v interface{}) string {
		a, _ := json.MarshalIndent(v, "", "  ")
		return string(a)
	},
	"splitArray": func(sep string, s string) []string {
		return strings.Split(s, sep)
	},
	"first": func(a []string) string {
		return a[0]
	},
	"last": func(a []string) string {
		return a[len(a)-1]
	},
	"upperFirst": func(s string) string {
		return strings.ToUpper(s[:1]) + s[1:]
	},
	"lowerFirst": func(s string) string {
		return strings.ToLower(s[:1]) + s[1:]
	},
	"camelCase": func(s string) string {
		if len(s) > 1 {
			return xstrings.ToCamelCase(s)
		}

		return strings.ToUpper(s[:1])
	},
	"lowerCamelCase": func(s string) string {
		if len(s) > 1 {
			s = xstrings.ToCamelCase(s)
		}

		return strings.ToLower(s[:1]) + s[1:]
	},
	"snakeCase": func(s string) string {
		return xstrings.ToSnakeCase(s)
	},
	"kebabCase": func(s string) string {
		return strings.Replace(xstrings.ToSnakeCase(s), "_", "-", -1)
	},
	"getMessageType":  getMessageType,
	"isFieldMessage":  isFieldMessage,
	"isFieldRepeated": isFieldRepeated,
	"goType":          goType,
	"jsType":          jsType,
	"httpVerb":        httpVerb,
	"httpPath":        httpPath,
	"shortType":       shortType,
}

func init() {
	for k, v := range sprig.TxtFuncMap() {
		ProtoHelpersFuncMap[k] = v
	}
}

func getMessageType(f *descriptor.FileDescriptorProto, name string) *descriptor.DescriptorProto {
	// name is in the form .packageName.MessageTypeName.InnerMessageTypeName...
	// e.g. .article.ProductTag
	splits := strings.Split(name, ".")
	target := splits[len(splits)-1]
	for _, m := range f.MessageType {
		if target == *m.Name {
			return m
		}
	}

	return nil
}

func isFieldMessage(f *descriptor.FieldDescriptorProto) bool {
	if f.Type != nil && *f.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
		return true
	}

	return false
}

func isFieldRepeated(f *descriptor.FieldDescriptorProto) bool {
	if f.Type != nil && f.Label != nil && *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
		return true
	}

	return false
}

func goType(pkg string, f *descriptor.FieldDescriptorProto) string {
	switch *f.Type {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return "float64"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		return "float32"
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return "int64"
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		return "uint64"
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return "uint32"
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return "bool"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return "string"
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return fmt.Sprintf("[]*%s.%s", pkg, shortType(*f.TypeName))
		}
		return fmt.Sprintf("*%s.%s", pkg, shortType(*f.TypeName))
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return "byte"
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		return "uint32"
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		return fmt.Sprintf("*%s.%s", pkg, shortType(*f.TypeName))
	default:
		return "interface{}"
	}
}

func jsType(f *descriptor.FieldDescriptorProto) string {
	template := "%s"
	if isFieldRepeated(f) == true {
		template = "Array<%s>"
	}

	switch *f.Type {
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		return fmt.Sprintf(template, shortType(*f.TypeName))
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT64:
		return fmt.Sprintf(template, "number")
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return fmt.Sprintf(template, "boolean")
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return fmt.Sprintf(template, "Array<number>")
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return fmt.Sprintf(template, "string")
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		return fmt.Sprintf(template, "Object")
	default:
		return fmt.Sprintf(template, "any")
	}
}

func shortType(s string) string {
	t := strings.Split(s, ".")
	return t[len(t)-1]
}

func httpPath(m *descriptor.MethodDescriptorProto) string {

	ext, err := proto.GetExtension(m.Options, options.E_Http)
	if err != nil {
		return err.Error()
	}
	opts, ok := ext.(*options.HttpRule)
	if !ok {
		return fmt.Sprintf("extension is %T; want an HttpRule", ext)
	}

	switch t := opts.Pattern.(type) {
	default:
		return ""
	case *options.HttpRule_Get:
		return t.Get
	case *options.HttpRule_Post:
		return t.Post
	case *options.HttpRule_Put:
		return t.Put
	case *options.HttpRule_Delete:
		return t.Delete
	case *options.HttpRule_Patch:
		return t.Patch
	case *options.HttpRule_Custom:
		return t.Custom.Path
	}
}

func httpVerb(m *descriptor.MethodDescriptorProto) string {

	ext, err := proto.GetExtension(m.Options, options.E_Http)
	if err != nil {
		return err.Error()
	}
	opts, ok := ext.(*options.HttpRule)
	if !ok {
		return fmt.Sprintf("extension is %T; want an HttpRule", ext)
	}

	switch t := opts.Pattern.(type) {
	default:
		return ""
	case *options.HttpRule_Get:
		return "GET"
	case *options.HttpRule_Post:
		return "POST"
	case *options.HttpRule_Put:
		return "PUT"
	case *options.HttpRule_Delete:
		return "DELETE"
	case *options.HttpRule_Patch:
		return "PATCH"
	case *options.HttpRule_Custom:
		return t.Custom.Kind
	}
}
