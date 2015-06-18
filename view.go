package goninja

import (
	"strings"
//	"html"
	"html/template"
)






var views_path = CURRENT_DIR + "/app/views/"
var layout_path = CURRENT_DIR + "/app/views/layouts/base.tmpl"



var FuncMap = template.FuncMap{
	"whatEver": func(s string) string {return s},
}


type View struct {
	Name string
	Path string
	C *Controller
}


func (v *View) RenderView() {


	tmpl := template.New("base.tmpl").Funcs(FuncMap)
	tmpl.ParseFiles(layout_path, v.TemplatePath(v.C.Action), v.TemplatePath("header"))


	if err == nil {
		tmpl.Execute(v.C.Writer, nil)
	} else {
		LOGGER.Println(err)
		v.C.Writer.Write([]byte(v.TemplateNotFound()))
	}

}



func (v *View) TemplateNotFound() string {
	return "Template with name \"" + strings.ToLower(v.C.Action) + "\" cound not be found or contains errors"
}

func (v *View) TemplatePath(name string) string {
	return views_path + strings.ToLower(v.Name) + "/" + strings.ToLower(name) + ".html"
}

