package goninja

import (
	"strings"
	"html"
)

//File with some basic helpers

func InitHelpers() {
	FuncMap["css_asset_path"] = func(path string) string {
		name := path
		if strings.HasSuffix(path, ".css") {
			name = "/assets/css/" + path
		} else {
			name = "/assets/css/" + path + ".css"
		}
		return stylesheet_tag(name)
	}
}




func stylesheet_tag(filename string) string {
	return html.UnescapeString(`<link rel="stylesheet" type="text/css" href="` + filename + `">`)
}
