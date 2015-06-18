package goninja

import (
	"strings"
	"html"
)

//File with some basic helpers
//TODO i have no idea how to organize this piece of code, since my experience with golang less than 3 weeks
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

func javascript_tag(filename string) string {
	return html.UnescapeString(`<script src="` + filename + `">`)
}
