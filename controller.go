package goninja

import (
    "net/http"
    "time"
    "strings"
    "encoding/json"
    "reflect"
)

type Controller struct {
    Name string
    Action string
    Format string
    Writer http.ResponseWriter
    Request *http.Request
    Views []string
}

const JSON_FORMAT string = "application/json"
const HTML_FORMAT string = "text/html"
const XML_FORMAT string = "text/xml"

type Response struct {
    Content string
}

var App_controllers map[string]interface{} = make(map[string]interface{})

func (c *Controller) Render(data interface{}, params ...interface{}) Response {
    c.Params();
    timeStarted := time.Now()
    view := c.createView();
    requestFormat := c.RequestFormat()
    if (requestFormat == HTML_FORMAT) {
        //    if viewParams := params["views"]; viewParams != nil {
        //        view.RenderView(viewParams)
        //
        //    } else {
        view.RenderView(data)
        //    }
    } else if (requestFormat == JSON_FORMAT) {
        //view.RenderJson(data)
    }
    extractParams(params)

    diff := time.Since(timeStarted).String()
    LOGGER.Printf("Views processed in %s", diff)
    return Response{}
}

func (c *Controller) RenderJson(jsonPayload interface{}) Response {

    jsonString, err := json.Marshal(jsonPayload);
    LOGGER.Println(jsonString);
    if err == nil {
        view := c.createView();
        view.RenderJson(jsonString)
    } else {
        LOGGER.Println(err);
    }
    return Response{}
}

func (c *Controller) createView() View {
    return View{
        Name: c.Name,
        C: c,
    }
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

func (c *Controller) RequestFormat() string {
    parsedFromHeader := c.Request.Header["Content-Type"]
    LOGGER.Println(reflect.TypeOf(parsedFromHeader));
    return "123"
}

func LaunchController(name string) (s interface{}, r bool) {
    if ctrl, err:= App_controllers[name]; err {
        return ctrl, err
    } else {
        return 0, false
    }
}

