package goninja

import "os"
import (
	"strings"
	"html/template"
)

var CURRENT_DIR, err = os.Getwd()



var VIEWS_PATH = CURRENT_DIR + "/app/views/"



type View struct {
	Name string
	Path string
	C *Controller
}


func (v *View) RenderView() {
	tmpl, err := template.ParseFiles(v.TemplatePath())
	if err ==nil {
		tmpl.Execute(v.C.Writer, nil)
	} else {
		v.C.Writer.Write([]byte(v.TemplateNotFound()))
	}

}

func (v *View) TemplateNotFound() string {
	return "Template with name \"" + strings.ToLower(v.C.Action) + "\" cound not be found"
}

func (v *View) TemplatePath() string {
	return VIEWS_PATH + strings.ToLower(v.C.Name) + "/" + strings.ToLower(v.C.Action) + ".html"
}
