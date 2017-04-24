package example

//go:generate protoc --go_out=./gen/         example.proto
//go:generate protoc --gotemplate_out=./gen/ example.proto
