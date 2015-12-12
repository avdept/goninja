package goninja

import (
    "net/http"
    "time"
    "strings"
)

type Controller struct {
    Name string
    Action string
    Format string
    Writer http.ResponseWriter
    Request *http.Request
    Views []string
}


type Response struct {
    Content string
}

var App_controllers map[string]interface{} = make(map[string]interface{})

func (c *Controller) Render(data interface{}, params ...interface{}) Response {
    c.Params();
    timeStarted := time.Now()
    view := View{
        Name: c.Name,
        C: c,
    }
    extractParams(params)
//    if viewParams := params["views"]; viewParams != nil {
//        view.RenderView(viewParams)
//
//    } else {
        view.RenderView(data)
//    }
    diff := time.Since(timeStarted).String()
    LOGGER.Printf("Views processed in %s", diff)
    return Response{}
}

func (c *Controller) Redirect(url string) Response {
//    LOGGER.Println("1231321312321") TODO redirect notice
    http.Redirect(c.Writer, c.Request, url, http.StatusMovedPermanently)
    return Response{}
}

func (c *Controller) Params() map[string]interface{} {
    result := make(map[string]interface{})
    c.Request.ParseForm()
    for k, v := range c.Request.Form{
        if len(v) > 1 {
            result[k] = v
        } else {
            result[k] = strings.Join(v, "")
        }
    }
    return result
}

func CreateControllers(name string, c interface{}) {
    App_controllers[name] = c
}

func extractParams(params ...interface{}) {
}

func LaunchController(name string) (s interface{}, r bool) {
    if ctrl, err:= App_controllers[name]; err {
        return ctrl, err
    } else {
        return 0, false
    }
}

