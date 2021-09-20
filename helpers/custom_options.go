package pgghelpers

import (
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto" // nolint:staticcheck
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoimpl"
)

var protoregistryMutex sync.Mutex

// getExtension returns the value of an extension.
//
// If the extension with the given ID is not already registered, it will be made up.
func getExtension(extendedMessage proto.Message, extendedType proto.Message, fieldID int32, fieldType interface{}) (interface{}, error) {
	// To prevent concurrent map read/write while querying the registry and registring new extensions, request a lock.
	protoregistryMutex.Lock()
	defer protoregistryMutex.Unlock()

	// Query the registry for the given Message and field ID.
	eds := make(map[int32]*protoimpl.ExtensionInfo)
	protoregistry.GlobalTypes.RangeExtensionsByMessage(protoimpl.X.MessageDescriptorOf(extendedMessage).FullName(), func(xt protoreflect.ExtensionType) bool {
		if xd, ok := xt.(*protoimpl.ExtensionInfo); ok {
			eds[int32(xt.TypeDescriptor().Number())] = xd
		}

		return true
	})
	extensionInfo := eds[fieldID]

	// Create the extension, if it was not yet found.
	if extensionInfo == nil {
		// Infer the struct tag type from field type.
		tagType := "varint"
		if _, ok := fieldType.(*string); ok {
			tagType = "bytes"
		}

		extensionInfo = &proto.ExtensionDesc{
			ExtendedType:  extendedType,
			ExtensionType: fieldType,
			Field:         fieldID,
			Tag:           fmt.Sprintf("%s,%d", tagType, fieldID),
			Name:          fmt.Sprintf("%d", fieldID),
		}
		err := protoregistry.GlobalTypes.RegisterExtension(extensionInfo)
		if err != nil {
			return nil, fmt.Errorf("error registering extension: %w", err)
		}
	}

	return proto.GetExtension(extendedMessage, extensionInfo)
}

// stringMethodOptionsExtension extracts method options of a string type.
// To define your own extensions see:
// https://developers.google.com/protocol-buffers/docs/proto#customoptions
// Typically the fieldID of private extensions should be in the range:
// 50000-99999
func stringMethodOptionsExtension(fieldID int32, f *descriptor.MethodDescriptorProto) string {
	if f == nil || f.Options == nil {
		return ""
	}

	var extendedType *descriptor.MethodOptions
	var fieldType *string

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return ""
	}

	if str, ok := ext.(*string); ok {
		return *str
	}

	return ""
}

// boolMethodOptionsExtension extracts method options of a boolean type.
func boolMethodOptionsExtension(fieldID int32, f *descriptor.MethodDescriptorProto) bool {
	if f == nil || f.Options == nil {
		return false
	}

	var extendedType *descriptor.MethodOptions
	var fieldType *bool

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return false
	}

	if b, ok := ext.(*bool); ok {
		return *b
	}

	return false
}

// stringFileOptionsExtension extracts file options of a string type.
// To define your own extensions see:
// https://developers.google.com/protocol-buffers/docs/proto#customoptions
// Typically the fieldID of private extensions should be in the range:
// 50000-99999
func stringFileOptionsExtension(fieldID int32, f *descriptor.FileDescriptorProto) string {
	if f == nil || f.Options == nil {
		return ""
	}

	var extendedType *descriptor.FileOptions
	var fieldType *string

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return ""
	}

	if str, ok := ext.(*string); ok {
		return *str
	}

	return ""
}

func stringFieldExtension(fieldID int32, f *descriptor.FieldDescriptorProto) string {
	if f == nil || f.Options == nil {
		return ""
	}

	var extendedType *descriptor.FieldOptions
	var fieldType *string

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return ""
	}

	str, ok := ext.(*string)
	if !ok {
		return ""
	}

	return *str
}

func int64FieldExtension(fieldID int32, f *descriptor.FieldDescriptorProto) int64 {
	if f == nil || f.Options == nil {
		return 0
	}

	var extendedType *descriptor.FieldOptions
	var fieldType *int64

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return 0
	}

	i, ok := ext.(*int64)
	if !ok {
		return 0
	}

	return *i
}

func int64MessageExtension(fieldID int32, f *descriptor.DescriptorProto) int64 {
	if f == nil || f.Options == nil {
		return 0
	}

	var extendedType *descriptor.MessageOptions
	var fieldType *int64

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return 0
	}

	i, ok := ext.(*int64)
	if !ok {
		return 0
	}

	return *i
}

func stringMessageExtension(fieldID int32, f *descriptor.DescriptorProto) string {
	if f == nil || f.Options == nil {
		return ""
	}

	var extendedType *descriptor.MessageOptions
	var fieldType *string

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return ""
	}

	str, ok := ext.(*string)
	if !ok {
		return ""
	}

	return *str
}

func boolFieldExtension(fieldID int32, f *descriptor.FieldDescriptorProto) bool {
	if f == nil || f.Options == nil {
		return false
	}

	var extendedType *descriptor.FieldOptions
	var fieldType *bool

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return false
	}

	b, ok := ext.(*bool)
	if !ok {
		return false
	}

	return *b
}

func boolMessageExtension(fieldID int32, f *descriptor.DescriptorProto) bool {
	if f == nil || f.Options == nil {
		return false
	}
	var extendedType *descriptor.MessageOptions
	var fieldType *bool

	ext, err := getExtension(f.Options, extendedType, fieldID, fieldType)
	if err != nil {
		return false
	}

	b, ok := ext.(*bool)
	if !ok {
		return false
	}

	return *b
}
