package pgghelpers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	ggdescriptor "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	"github.com/huandu/xstrings"
	options "google.golang.org/genproto/googleapis/api/annotations"
)

var jsReservedRe = regexp.MustCompile(`(^|[^A-Za-z])(do|if|in|for|let|new|try|var|case|else|enum|eval|false|null|this|true|void|with|break|catch|class|const|super|throw|while|yield|delete|export|import|public|return|static|switch|typeof|default|extends|finally|package|private|continue|debugger|function|arguments|interface|protected|implements|instanceof)($|[^A-Za-z])`)

var (
	registry *ggdescriptor.Registry // some helpers need access to registry
)

var ProtoHelpersFuncMap = template.FuncMap{
	"string": func(i interface {
		String() string
	}) string {
		return i.String()
	},
	"json": func(v interface{}) string {
		a, err := json.Marshal(v)
		if err != nil {
			return err.Error()
		}
		return string(a)
	},
	"prettyjson": func(v interface{}) string {
		a, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return err.Error()
		}
		return string(a)
	},
	"splitArray": func(sep string, s string) []interface{} {
		var r []interface{}
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
	"concat": func(a string, b ...string) string {
		return strings.Join(append([]string{a}, b...), "")
	},
	"join": func(sep string, a ...string) string {
		return strings.Join(a, sep)
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
	"contains": func(sub, s string) bool {
		return strings.Contains(s, sub)
	},
	"trimstr": func(cutset, s string) string {
		return strings.Trim(s, cutset)
	},
	"index": func(array interface{}, i int32) interface{} {
		slice := reflect.ValueOf(array)
		if slice.Kind() != reflect.Slice {
			panic("Error in index(): given a non-slice type")
		}
		if i < 0 || int(i) >= slice.Len() {
			panic("Error in index(): index out of bounds")
		}
		return slice.Index(int(i)).Interface()
	},
	"add": func(a int, b int) int {
		return a + b
	},
	"subtract": func(a int, b int) int {
		return a - b
	},
	"multiply": func(a int, b int) int {
		return a * b
	},
	"divide": func(a int, b int) int {
		if b == 0 {
			panic("psssst ... little help here ... you cannot divide by 0")
		}
		return a / b
	},

	"snakeCase":                    xstrings.ToSnakeCase,
	"getProtoFile":                 getProtoFile,
	"getMessageType":               getMessageType,
	"getEnumValue":                 getEnumValue,
	"isFieldMessage":               isFieldMessage,
	"isFieldMessageTimeStamp":      isFieldMessageTimeStamp,
	"isFieldRepeated":              isFieldRepeated,
	"haskellType":                  haskellType,
	"goType":                       goType,
	"goZeroValue":                  goZeroValue,
	"goTypeWithPackage":            goTypeWithPackage,
	"goTypeWithGoPackage":          goTypeWithGoPackage,
	"jsType":                       jsType,
	"jsSuffixReserved":             jsSuffixReservedKeyword,
	"namespacedFlowType":           namespacedFlowType,
	"httpVerb":                     httpVerb,
	"httpPath":                     httpPath,
	"httpPathsAdditionalBindings":  httpPathsAdditionalBindings,
	"httpBody":                     httpBody,
	"shortType":                    shortType,
	"urlHasVarsFromMessage":        urlHasVarsFromMessage,
	"lowerGoNormalize":             lowerGoNormalize,
	"goNormalize":                  goNormalize,
	"leadingComment":               leadingComment,
	"trailingComment":              trailingComment,
	"leadingDetachedComments":      leadingDetachedComments,
	"stringFieldExtension":         stringFieldExtension,
	"stringMethodOptionsExtension": stringMethodOptionsExtension,
	"boolMethodOptionsExtension":   boolMethodOptionsExtension,
	"boolFieldExtension":           boolFieldExtension,
	"isFieldMap":                   isFieldMap,
	"fieldMapKeyType":              fieldMapKeyType,
	"fieldMapValueType":            fieldMapValueType,
	"replaceDict":                  replaceDict,
	"setStore":                     setStore,
	"getStore":                     getStore,
}

var pathMap map[interface{}]*descriptor.SourceCodeInfo_Location

var store = newStore()

// Utility to store some vars across multiple scope
type globalStore struct {
	store map[string]interface{}
	mu    sync.Mutex
}

func newStore() *globalStore {
	return &globalStore{
		store: make(map[string]interface{}),
	}
}

func (s *globalStore) getData(key string) interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	if v, ok := s.store[key]; ok {
		return v
	}

	return false
}

func (s *globalStore) setData(key string, o interface{}) {
	s.mu.Lock()
	s.store[key] = o
	s.mu.Unlock()
}

func setStore(key string, o interface{}) string {
	store.setData(key, o)
	return ""
}

func getStore(key string) interface{} {
	return store.getData(key)
}

func SetRegistry(reg *ggdescriptor.Registry) {
	registry = reg
}

func InitPathMap(file *descriptor.FileDescriptorProto) {
	pathMap = make(map[interface{}]*descriptor.SourceCodeInfo_Location)
	addToPathMap(file.GetSourceCodeInfo(), file, []int32{})
}

func InitPathMaps(files []*descriptor.FileDescriptorProto) {
	pathMap = make(map[interface{}]*descriptor.SourceCodeInfo_Location)
	for _, file := range files {
		addToPathMap(file.GetSourceCodeInfo(), file, []int32{})
	}
}

// addToPathMap traverses through the AST adding SourceCodeInfo_Location entries to the pathMap.
// Since the AST is a tree, the recursion finishes once it has gone through all the nodes.
func addToPathMap(info *descriptor.SourceCodeInfo, i interface{}, path []int32) {
	loc := findLoc(info, path)
	if loc != nil {
		pathMap[i] = loc
	}
	switch d := i.(type) {
	case *descriptor.FileDescriptorProto:
		for index, descriptor := range d.MessageType {
			addToPathMap(info, descriptor, newPath(path, 4, index))
		}
		for index, descriptor := range d.EnumType {
			addToPathMap(info, descriptor, newPath(path, 5, index))
		}
		for index, descriptor := range d.Service {
			addToPathMap(info, descriptor, newPath(path, 6, index))
		}
	case *descriptor.DescriptorProto:
		for index, descriptor := range d.Field {
			addToPathMap(info, descriptor, newPath(path, 2, index))
		}
		for index, descriptor := range d.NestedType {
			addToPathMap(info, descriptor, newPath(path, 3, index))
		}
		for index, descriptor := range d.EnumType {
			addToPathMap(info, descriptor, newPath(path, 4, index))
		}
	case *descriptor.EnumDescriptorProto:
		for index, descriptor := range d.Value {
			addToPathMap(info, descriptor, newPath(path, 2, index))
		}
	case *descriptor.ServiceDescriptorProto:
		for index, descriptor := range d.Method {
			addToPathMap(info, descriptor, newPath(path, 2, index))
		}
	}
}

func newPath(base []int32, field int32, index int) []int32 {
	p := append([]int32{}, base...)
	p = append(p, field, int32(index))
	return p
}

func findLoc(info *descriptor.SourceCodeInfo, path []int32) *descriptor.SourceCodeInfo_Location {
	for _, loc := range info.GetLocation() {
		if samePath(loc.Path, path) {
			return loc
		}
	}
	return nil
}

func samePath(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, p := range a {
		if p != b[i] {
			return false
		}
	}
	return true
}

func findSourceInfoLocation(i interface{}) *descriptor.SourceCodeInfo_Location {
	if pathMap == nil {
		return nil
	}
	return pathMap[i]
}

func leadingComment(i interface{}) string {
	loc := pathMap[i]
	return loc.GetLeadingComments()
}
func trailingComment(i interface{}) string {
	loc := pathMap[i]
	return loc.GetTrailingComments()
}
func leadingDetachedComments(i interface{}) []string {
	loc := pathMap[i]
	return loc.GetLeadingDetachedComments()
}

// stringMethodOptionsExtension extracts method options of a string type.
// To define your own extensions see:
// https://developers.google.com/protocol-buffers/docs/proto#customoptions
// Typically the fieldID of private extensions should be in the range:
// 50000-99999
func stringMethodOptionsExtension(fieldID int32, f *descriptor.MethodDescriptorProto) string {
	if f == nil {
		return ""
	}
	if f.Options == nil {
		return ""
	}
	var extendedType *descriptor.MethodOptions
	var extensionType *string

	eds := proto.RegisteredExtensions(f.Options)
	if eds[fieldID] == nil {
		ed := &proto.ExtensionDesc{
			ExtendedType:  extendedType,
			ExtensionType: extensionType,
			Field:         fieldID,
			Tag:           fmt.Sprintf("bytes,%d", fieldID),
		}
		proto.RegisterExtension(ed)
		eds = proto.RegisteredExtensions(f.Options)
	}

	ext, err := proto.GetExtension(f.Options, eds[fieldID])
	if err != nil {
		return ""
	}

	str, ok := ext.(*string)
	if !ok {
		return ""
	}

	return *str
}

func stringFieldExtension(fieldID int32, f *descriptor.FieldDescriptorProto) string {
	if f == nil {
		return ""
	}
	if f.Options == nil {
		return ""
	}
	var extendedType *descriptor.FieldOptions
	var extensionType *string

	eds := proto.RegisteredExtensions(f.Options)
	if eds[fieldID] == nil {
		ed := &proto.ExtensionDesc{
			ExtendedType:  extendedType,
			ExtensionType: extensionType,
			Field:         fieldID,
			Tag:           fmt.Sprintf("bytes,%d", fieldID),
		}
		proto.RegisterExtension(ed)
		eds = proto.RegisteredExtensions(f.Options)
	}

	ext, err := proto.GetExtension(f.Options, eds[fieldID])
	if err != nil {
		return ""
	}

	str, ok := ext.(*string)
	if !ok {
		return ""
	}

	return *str
}

func boolMethodOptionsExtension(fieldID int32, f *descriptor.MethodDescriptorProto) bool {
	if f == nil {
		return false
	}
	if f.Options == nil {
		return false
	}
	var extendedType *descriptor.MethodOptions
	var extensionType *bool

	eds := proto.RegisteredExtensions(f.Options)
	if eds[fieldID] == nil {
		ed := &proto.ExtensionDesc{
			ExtendedType:  extendedType,
			ExtensionType: extensionType,
			Field:         fieldID,
			Tag:           fmt.Sprintf("bytes,%d", fieldID),
		}
		proto.RegisterExtension(ed)
		eds = proto.RegisteredExtensions(f.Options)
	}

	ext, err := proto.GetExtension(f.Options, eds[fieldID])
	if err != nil {
		return false
	}

	b, ok := ext.(*bool)
	if !ok {
		return false
	}

	return *b
}

func boolFieldExtension(fieldID int32, f *descriptor.FieldDescriptorProto) bool {
	if f == nil {
		return false
	}
	if f.Options == nil {
		return false
	}
	var extendedType *descriptor.FieldOptions
	var extensionType *bool

	eds := proto.RegisteredExtensions(f.Options)
	if eds[fieldID] == nil {
		ed := &proto.ExtensionDesc{
			ExtendedType:  extendedType,
			ExtensionType: extensionType,
			Field:         fieldID,
			Tag:           fmt.Sprintf("varint,%d", fieldID),
		}
		proto.RegisterExtension(ed)
		eds = proto.RegisteredExtensions(f.Options)
	}

	ext, err := proto.GetExtension(f.Options, eds[fieldID])
	if err != nil {
		return false
	}

	b, ok := ext.(*bool)
	if !ok {
		return false
	}

	return *b
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

func isFieldMessageTimeStamp(f *descriptor.FieldDescriptorProto) bool {
	if f.Type != nil && *f.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
		if strings.Compare(*f.TypeName, ".google.protobuf.Timestamp") == 0 {
			return true
		}
	}
	return false
}

func isFieldMessage(f *descriptor.FieldDescriptorProto) bool {
	if f.Type != nil && *f.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
		return true
	}

	return false
}

func isFieldRepeated(f *descriptor.FieldDescriptorProto) bool {
	if f == nil {
		return false
	}
	if f.Type != nil && f.Label != nil && *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
		return true
	}

	return false
}

func isFieldMap(f *descriptor.FieldDescriptorProto, m *descriptor.DescriptorProto) bool {
	if f.TypeName == nil {
		return false
	}

	shortName := shortType(*f.TypeName)
	var nt *descriptor.DescriptorProto
	for _, t := range m.NestedType {
		if *t.Name == shortName {
			nt = t
			break
		}
	}

	if nt == nil {
		return false
	}

	for _, f := range nt.Field {
		switch *f.Name {
		case "key":
			if *f.Number != 1 {
				return false
			}
		case "value":
			if *f.Number != 2 {
				return false
			}
		default:
			return false
		}
	}

	return true
}

func fieldMapKeyType(f *descriptor.FieldDescriptorProto, m *descriptor.DescriptorProto) *descriptor.FieldDescriptorProto {
	if f.TypeName == nil {
		return nil
	}

	shortName := shortType(*f.TypeName)
	var nt *descriptor.DescriptorProto
	for _, t := range m.NestedType {
		if *t.Name == shortName {
			nt = t
			break
		}
	}

	if nt == nil {
		return nil
	}

	for _, f := range nt.Field {
		if *f.Name == "key" {
			return f
		}
	}

	return nil

}

func fieldMapValueType(f *descriptor.FieldDescriptorProto, m *descriptor.DescriptorProto) *descriptor.FieldDescriptorProto {
	if f.TypeName == nil {
		return nil
	}

	shortName := shortType(*f.TypeName)
	var nt *descriptor.DescriptorProto
	for _, t := range m.NestedType {
		if *t.Name == shortName {
			nt = t
			break
		}
	}

	if nt == nil {
		return nil
	}

	for _, f := range nt.Field {
		if *f.Name == "value" {
			return f
		}
	}

	return nil

}

// goTypeWithGoPackage types the field MESSAGE and ENUM with the go_package name.
// This method is an evolution of goTypeWithPackage. It handles message embedded.
//
// example:
// ```proto
// message GetArticleResponse {
// 	Article article = 1;
// 	message Storage {
// 		  string code = 1;
// 	}
// 	repeated Storage storages = 2;
// }
// ```
// Then the type of `storages` is `GetArticleResponse_Storage` for the go language.
//
func goTypeWithGoPackage(p *descriptor.FileDescriptorProto, f *descriptor.FieldDescriptorProto) string {
	pkg := ""
	if *f.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE || *f.Type == descriptor.FieldDescriptorProto_TYPE_ENUM {
		if isTimestampPackage(*f.TypeName) {
			pkg = "timestamp"
		} else {
			pkg = *p.GetOptions().GoPackage
			if strings.Contains(*p.GetOptions().GoPackage, ";") {
				pkg = strings.Split(*p.GetOptions().GoPackage, ";")[1]
			}
		}
	}
	return goTypeWithEmbedded(pkg, f, p)
}

// Warning does not handle message embedded like goTypeWithGoPackage does.
func goTypeWithPackage(f *descriptor.FieldDescriptorProto) string {
	pkg := ""
	if *f.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE || *f.Type == descriptor.FieldDescriptorProto_TYPE_ENUM {
		if isTimestampPackage(*f.TypeName) {
			pkg = "timestamp"
		} else {
			pkg = getPackageTypeName(*f.TypeName)
		}
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

func goTypeWithEmbedded(pkg string, f *descriptor.FieldDescriptorProto, p *descriptor.FileDescriptorProto) string {
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
		name := *f.TypeName
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			fieldPackage := strings.Split(*f.TypeName, ".")
			filePackage := strings.Split(*p.Package, ".")
			// check if we are working with a message embedded.
			if len(fieldPackage) > 1 && len(fieldPackage)+1 > len(filePackage)+1 {
				name = strings.Join(fieldPackage[len(filePackage)+1:], "_")
			}

			return fmt.Sprintf("[]*%s%s", pkg, shortType(name))
		}
		return fmt.Sprintf("*%s%s", pkg, shortType(name))
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return "[]byte"
		}
		return "byte"
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		name := *f.TypeName
		fieldPackage := strings.Split(*f.TypeName, ".")
		filePackage := strings.Split(*p.Package, ".")
		// check if we are working with a message embedded.
		if len(fieldPackage) > 1 && len(fieldPackage)+1 > len(filePackage)+1 {
			name = strings.Join(fieldPackage[len(filePackage)+1:], "_")
		}
		return fmt.Sprintf("*%s%s", pkg, shortType(name))
	default:
		return "interface{}"
	}
}

//Deprecated. Instead use goTypeWithEmbedded
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

func goZeroValue(f *descriptor.FieldDescriptorProto) string {
	const nilString = "nil"
	if *f.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
		return nilString
	}
	switch *f.Type {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return "0.0"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		return "0.0"
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return "0"
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		return "0"
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return "0"
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		return "0"
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return "false"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return "\"\""
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		return nilString
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return "0"
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		return nilString
	default:
		return nilString
	}
}

func jsType(f *descriptor.FieldDescriptorProto) string {
	template := "%s"
	if isFieldRepeated(f) {
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

func isTimestampPackage(s string) bool {
	var isTimestampPackage bool
	if strings.Compare(s, ".google.protobuf.Timestamp") == 0 {
		isTimestampPackage = true
	}
	return isTimestampPackage
}

func getPackageTypeName(s string) string {
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

func httpPathsAdditionalBindings(m *descriptor.MethodDescriptorProto) []string {
	ext, err := proto.GetExtension(m.Options, options.E_Http)
	if err != nil {
		panic(err.Error())
	}
	opts, ok := ext.(*options.HttpRule)
	if !ok {
		panic(fmt.Sprintf("extension is %T; want an HttpRule", ext))
	}

	var httpPaths []string
	var optsAdditionalBindings = opts.GetAdditionalBindings()
	for _, optAdditionalBindings := range optsAdditionalBindings {
		switch t := optAdditionalBindings.Pattern.(type) {
		case *options.HttpRule_Get:
			httpPaths = append(httpPaths, t.Get)
		case *options.HttpRule_Post:
			httpPaths = append(httpPaths, t.Post)
		case *options.HttpRule_Put:
			httpPaths = append(httpPaths, t.Put)
		case *options.HttpRule_Delete:
			httpPaths = append(httpPaths, t.Delete)
		case *options.HttpRule_Patch:
			httpPaths = append(httpPaths, t.Patch)
		case *options.HttpRule_Custom:
			httpPaths = append(httpPaths, t.Custom.Path)
		default:
			// nothing
		}
	}

	return httpPaths
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

func httpBody(m *descriptor.MethodDescriptorProto) string {

	ext, err := proto.GetExtension(m.Options, options.E_Http)
	if err != nil {
		return err.Error()
	}
	opts, ok := ext.(*options.HttpRule)
	if !ok {
		return fmt.Sprintf("extension is %T; want an HttpRule", ext)
	}
	return opts.Body
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

// lowerGoNormalize takes a string and applies formatting
// rules to conform to Golang convention. It applies a camel
// case filter, lowers the first character and formats fields
// with `id` to `ID`.
func lowerGoNormalize(s string) string {
	fmtd := xstrings.ToCamelCase(s)
	fmtd = xstrings.FirstRuneToLower(fmtd)
	return formatID(s, fmtd)
}

// goNormalize takes a string and applies formatting rules
// to conform to Golang convention. It applies a camel case
// filter and formats fields with `id` to `ID`.
func goNormalize(s string) string {
	fmtd := xstrings.ToCamelCase(s)
	return formatID(s, fmtd)
}

// formatID takes a base string alonsgide a formatted string.
// It acts as a transformation filter for fields containing
// `id` in order to conform to Golang convention.
func formatID(base string, formatted string) string {
	if formatted == "" {
		return formatted
	}
	switch {
	case base == "id":
		// id -> ID
		return "ID"
	case strings.HasPrefix(base, "id_"):
		// id_some -> IDSome
		return "ID" + formatted[2:]
	case strings.HasSuffix(base, "_id"):
		// some_id -> SomeID
		return formatted[:len(formatted)-2] + "ID"
	case strings.HasSuffix(base, "_ids"):
		// some_ids -> SomeIDs
		return formatted[:len(formatted)-3] + "IDs"
	}
	return formatted
}

func replaceDict(src string, dict map[string]interface{}) string {
	for old, v := range dict {
		new, ok := v.(string)
		if !ok {
			continue
		}
		src = strings.Replace(src, old, new, -1)
	}
	return src
}
