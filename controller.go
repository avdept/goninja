package goninja


import "net/http"
//import "reflect"

type Controller struct {
    Name string
    Writer http.ResponseWriter
    Request *http.Request
}


type Response struct {
    Content string
}


var App_controllers map[string]interface{} = make(map[string]interface{})



func CreateControllers(name string, c interface{}) {
    App_controllers[name] = c
}

func LaunchController(name string) interface{} {
    return App_controllers[name]
}

