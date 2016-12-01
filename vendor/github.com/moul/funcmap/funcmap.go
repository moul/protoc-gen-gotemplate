package funcmap

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"text/template"
)

var FuncMap = template.FuncMap{
	"json": func(v interface{}) string {
		a, _ := json.Marshal(v)
		return string(a)
	},
	"prettyjson": func(v interface{}) string {
		a, _ := json.MarshalIndent(v, "", "  ")
		return string(a)
	},
	// yaml
	// xml
	// toml
	"split": strings.Split,
	"join":  strings.Join,
	"title": strings.Title,
	"unexport": func(input string) string {
		return fmt.Sprintf("%s%s", strings.ToLower(input[0:1]), input[1:])
	},
	"lower": strings.ToLower,
	"upper": strings.ToUpper,
	"rev": func(v interface{}) string {
		runes := []rune(v.(string))
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	},
	"int": func(v interface{}) string {
		a, err := strconv.Atoi(v.(string))
		if err != nil {
			return fmt.Sprintf("%v", v)
		}
		return fmt.Sprintf("%d", a)
	},
}
