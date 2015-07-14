package goninja


import "net/http"
import (
    "time"
)



//import "reflect"

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

func (c *Controller) Render(params ...interface{}) Response {
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
        view.RenderView()
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



var App_controllers map[string]interface{} = make(map[string]interface{})



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

