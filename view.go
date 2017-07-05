package goninja

import (
	"strings"
	"html/template"
	"os"
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


func (v *View) RenderView(data interface{}) {

	path:= v.TemplatePath(v.C.Action)
	if !Exists(path){
		v.C.Writer.Write([]byte(v.TemplateNotFound()))
		return
	}

	tmpl := template.New("base.tmpl").Funcs(FuncMap)
	template.Must(tmpl.ParseFiles(layout_path, path))


	for _, view := range v.C.Views {
		LOGGER.Println(tmpl)
		tmpl, err =  tmpl.ParseFiles(v.TemplatePath(view))
	}

	if err == nil {
		LOGGER.Println(data)
		err = tmpl.Execute(v.C.Writer, data)
		if err != nil {
			v.C.Writer.Write([]byte(err.Error()))
		}
	} else {
//		LOGGER.Println(err)
		v.C.Writer.Write([]byte(v.TemplateNotFound()))
	}

}

func (v *View) RenderJson(jsonString []byte) {
	LOGGER.Println(jsonString);
	writer := v.C.Writer
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonString);
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (v *View) view_folder() string {
	path:= strings.Trim(v.Name, "Controller")
	return path
}



func (v *View) TemplateNotFound() string {
	return "Template with name \"" + strings.ToLower(v.C.Action) + "\" in \"views/" + v.view_folder() + "\" cound not be found or contains errors"
}

func (v *View) TemplatePath(name string) string {

	path := views_path + strings.ToLower(v.view_folder()) + "/" + strings.ToLower(name) + ".html"

	return path
}

