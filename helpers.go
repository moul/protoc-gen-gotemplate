package main

import (
	"encoding/json"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/huandu/xstrings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
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
		return xstrings.ToCamelCase(s)
	},
	"lowerCamelCase": func(s string) string {
		cc := xstrings.ToCamelCase(s)
		return strings.ToLower(cc[:1]) + cc[1:]
	},
	"snakeCase": func(s string) string {
		return xstrings.ToSnakeCase(s)
	},
	"kebabCase": func(s string) string {
		return strings.Replace(xstrings.ToSnakeCase(s), "_", "-", -1)
	},
	"getMessageType": getMessageType,
	"isFieldMessage": isFieldMessage,
}

func init() {
	for k, v := range sprig.TxtFuncMap() {
		ProtoHelpersFuncMap[k] = v
	}
}

func getMessageType(f *descriptor.FileDescriptorProto, name string) *descriptor.DescriptorProto {
	for _, m := range f.MessageType {
		// name usually contains the package name
		if strings.HasSuffix(name, *m.Name) {
			return m
		}
	}

	return nil
}

func isFieldMessage(f *descriptor.FieldDescriptorProto) bool {
	if f.Type != nil && *f.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE &&
		f.Label != nil && *f.Label != descriptor.FieldDescriptorProto_LABEL_REPEATED {
		return true
	}

	return false
}
