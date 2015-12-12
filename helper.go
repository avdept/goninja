package goninja

import (
	"strings"
	"html/template"
)

//File with some basic helpers
//TODO Refactor this, since I have no idea how to organize this piece of code
func InitHelpers() {
	FuncMap["css_asset_path"] = stylesheet_tag
	FuncMap["javascript_asset_path"] = javascript_tag

}




func stylesheet_tag(path string) template.HTML {
	name := path
	if strings.HasSuffix(path, ".css") {
		name = "/assets/css/" + path
	} else {
		name = "/assets/css/" + path + ".css"
	}
	return template.HTML(`<link rel="stylesheet" type="text/css" href="` + name + `">`)
}

func javascript_tag(filename string) template.HTML {
	name := filename
	if strings.HasSuffix(filename, ".js") {
		name = "/assets/js/" + filename
	} else {
		name = "/assets/js/" + filename + ".js"
	}
	return template.HTML(`<script src="` + name + `"></script>`)
}
