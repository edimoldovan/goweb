package router

import "html/template"

var defaultFuncs = template.FuncMap{
	"Title": func(ip interface{}) string {
		v, ok := ip.(string)
		if !ok || (ok && v == "") {
			return "Web app with Go std"
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
