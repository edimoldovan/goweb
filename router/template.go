package router

import "html/template"

// keeping for now as example
var defaultFuncs = template.FuncMap{
	"tmplFunctionExample": func(ip interface{}) string {
		v, ok := ip.(string)
		if !ok || (ok && v == "") {
			return "some text"
		}
		return v
	},
}
var templateFiles = []string{
	"./views/pages/default.html",
}

func tmplLayout(files ...string) []string {
	return append(templateFiles, files...)
}
