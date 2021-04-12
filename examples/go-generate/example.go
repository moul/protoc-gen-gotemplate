package example

//go:generate protoc --go_out=./gen/         example.proto
//go:generate protoc --gotemplate_out=template_dir=templates:./gen/ example.proto
