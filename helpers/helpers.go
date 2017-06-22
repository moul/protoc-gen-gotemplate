package pgghelpers

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	ggdescriptor "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	"github.com/huandu/xstrings"
	options "google.golang.org/genproto/googleapis/api/annotations"
)

var jsReservedRe *regexp.Regexp = regexp.MustCompile(`(^|[^A-Za-z])(do|if|in|for|let|new|try|var|case|else|enum|eval|false|null|this|true|void|with|break|catch|class|const|super|throw|while|yield|delete|export|import|public|return|static|switch|typeof|default|extends|finally|package|private|continue|debugger|function|arguments|interface|protected|implements|instanceof)($|[^A-Za-z])`)

var (
	registry *ggdescriptor.Registry // some helpers need access to registry
)

func SetRegistry(reg *ggdescriptor.Registry) {
	registry = reg
}

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
		var r []string
		t := strings.Split(s, sep)
		for i := range t {
			if t[i] != "" {
				r = append(r, t[i])
			}
		}
		return r
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
	"kebabCase": func(s string) string {
		return strings.Replace(xstrings.ToSnakeCase(s), "_", "-", -1)
	},
	"snakeCase":             xstrings.ToSnakeCase,
	"getProtoFile":          getProtoFile,
	"getMessageType":        getMessageType,
	"getEnumValue":          getEnumValue,
	"isFieldMessage":        isFieldMessage,
	"isFieldRepeated":       isFieldRepeated,
	"haskellType":           haskellType,
	"goType":                goType,
	"goTypeWithPackage":     goTypeWithPackage,
	"jsType":                jsType,
	"jsSuffixReserved":      jsSuffixReservedKeyword,
	"namespacedFlowType":    namespacedFlowType,
	"httpVerb":              httpVerb,
	"httpPath":              httpPath,
	"shortType":             shortType,
	"urlHasVarsFromMessage": urlHasVarsFromMessage,
}

func init() {
	for k, v := range sprig.TxtFuncMap() {
		ProtoHelpersFuncMap[k] = v
	}
}

func getProtoFile(name string) *ggdescriptor.File {
	if registry == nil {
		return nil
	}
	file, err := registry.LookupFile(name)
	if err != nil {
		panic(err)
	}
	return file
}

func getMessageType(f *descriptor.FileDescriptorProto, name string) *ggdescriptor.Message {
	if registry != nil {
		msg, err := registry.LookupMsg(".", name)
		if err != nil {
			panic(err)
		}
		return msg
	}

	// name is in the form .packageName.MessageTypeName.InnerMessageTypeName...
	// e.g. .article.ProductTag
	splits := strings.Split(name, ".")
	target := splits[len(splits)-1]
	for _, m := range f.MessageType {
		if target == *m.Name {
			return &ggdescriptor.Message{
				DescriptorProto: m,
			}
		}
	}
	return nil
}

func getEnumValue(f []*descriptor.EnumDescriptorProto, name string) []*descriptor.EnumValueDescriptorProto {
	for _, item := range f {
		if strings.EqualFold(*item.Name, name) {
			return item.GetValue()
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

func goTypeWithPackage(f *descriptor.FieldDescriptorProto) string {
	pkg := ""
	if *f.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
		pkg = getPackageTypeName(*f.TypeName)
	}
	return goType(pkg, f)
}

func haskellType(pkg string, f *descriptor.FieldDescriptorProto) string {
	switch *f.Type {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[Float]"
		}
		return "Float"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[Float]"
		}
		return "Float"
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[Int64]"
		}
		return "Int64"
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[Word]"
		}
		return "Word"
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[Int]"
		}
		return "Int"
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[Word]"
		}
		return "Word"
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[Bool]"
		}
		return "Bool"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[Text]"
		}
		return "Text"
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		if pkg != "" {
			pkg = pkg + "."
		}
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return fmt.Sprintf("[%s%s]", pkg, shortType(*f.TypeName))
		}
		return fmt.Sprintf("%s%s", pkg, shortType(*f.TypeName))
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[Word8]"
		}
		return "Word8"
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		return fmt.Sprintf("%s%s", pkg, shortType(*f.TypeName))
	default:
		return "Generic"
	}
}

func goType(pkg string, f *descriptor.FieldDescriptorProto) string {
	if pkg != "" {
		pkg = pkg + "."
	}
	switch *f.Type {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[]float64"
		}
		return "float64"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[]float32"
		}
		return "float32"
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[]int64"
		}
		return "int64"
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[]uint64"
		}
		return "uint64"
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[]int32"
		}
		return "int32"
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[]uint32"
		}
		return "uint32"
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[]bool"
		}
		return "bool"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[]string"
		}
		return "string"
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return fmt.Sprintf("[]*%s%s", pkg, shortType(*f.TypeName))
		}
		return fmt.Sprintf("*%s%s", pkg, shortType(*f.TypeName))
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[]byte"
		}
		return "byte"
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		return fmt.Sprintf("*%s%s", pkg, shortType(*f.TypeName))
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
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE,
		descriptor.FieldDescriptorProto_TYPE_ENUM:
		return fmt.Sprintf(template, namespacedFlowType(*f.TypeName))
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
		return fmt.Sprintf(template, "Uint8Array")
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return fmt.Sprintf(template, "string")
	default:
		return fmt.Sprintf(template, "any")
	}
}

func jsSuffixReservedKeyword(s string) string {
	return jsReservedRe.ReplaceAllString(s, "${1}${2}_${3}")
}

func getPackageTypeName(s string) string {
	if strings.Compare(s, ".google.protobuf.Timestamp") == 0 {
		return "timestamp"
	}
	if strings.Contains(s, ".") {
		return strings.Split(s, ".")[1]
	}
	return ""
}

func shortType(s string) string {
	t := strings.Split(s, ".")
	return t[len(t)-1]
}

func namespacedFlowType(s string) string {
	trimmed := strings.TrimLeft(s, ".")
	splitted := strings.Split(trimmed, ".")
	return strings.Join(splitted, "$")
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

func urlHasVarsFromMessage(path string, d *ggdescriptor.Message) bool {
	for _, field := range d.Field {
		if !isFieldMessage(field) {
			if strings.Contains(path, fmt.Sprintf("{%s}", *field.Name)) {
				return true
			}
		}
	}

	return false
}
