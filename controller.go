package goninja


import "net/http"
import "reflect"

type Controller struct {
    Name string
    Writer http.ResponseWriter
    Request *http.Request
}


type Response struct {
    content string
}


var App_controllers map[string]reflect.Type = make(map[string]reflect.Type)



func CreateControllers(c interface{}) {
    var t reflect.Type = reflect.TypeOf(c)
    var elem reflect.Type = t.Elem()
    App_controllers[elem.Name()] = t
    LOGGER.Println("Registered controller: ", elem.Name())
}

func LaunchController(name string) reflect.Type {
    return App_controllers[name]
}

