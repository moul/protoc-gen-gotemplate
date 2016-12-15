package main

import (
	"encoding/json"
	"text/template"

	"github.com/Masterminds/sprig"
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
	"first": func(a []interface{}) interface{} {
		return a[0]
	},
	"last": func(a []interface{}) interface{} {
		return a[len(a)-1]
	},
}

func init() {
	for k, v := range sprig.TxtFuncMap() {
		ProtoHelpersFuncMap[k] = v
	}
}
