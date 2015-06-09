package goninja

import (
	"fmt"
	"net/http"
	"reflect"
		// "strings"
	// "html/template"
	// "regexp"
)

type Params map[string]string

type Router struct {

	//slice with all routes
	routes []Route
	//hash with named route that matchers Route object
	named_routes map[string]*Route
}

type Route struct {
	method string
	pattern string
	action string
	controller string
}

func NewRouter() *Router {
	router := &Router{}
	return router
}

func (r *Router) AddRoute(method string, pattern string, action string, controller string) *Router {
	fmt.Println(controller)
	route := Route{method, pattern, action, controller}
//	CreateControllers((app_ctrl)(nil))
	r.routes = append(r.routes, route)
	LOGGER.Println("Router")
	return r
}

func (router *Router) match(r *http.Request) Route {
	fmt.Println(router.routes)
	return router.routes[0]
}

func  (router *Router) Handle(w http.ResponseWriter, r *http.Request) {
	LOGGER.Println("served from router")
	route := router.match(r)
		var t reflect.Type = LaunchController(route.controller)

		// create controller ptr .
		var appControllerPtr reflect.Value = reflect.New(t)
		fmt.Println(appControllerPtr)
		var appController reflect.Value = appControllerPtr.Elem()





		// Create and configure base controller
		var c  = &Controller{Request: r,	Writer: w,	Name: t.Name()	}


		//this should assign *goninja.Controller field in application controllers
		var controllerField reflect.Value = appController.FieldByName("Ctrl")
		fmt.Println(reflect.ValueOf(c).Kind())
		controllerField.Set(reflect.ValueOf(c))

		// Now call the action.
		// TODO: Figure out the arguments it expects, and try to bind parameters to
		// them.
	fmt.Println(controllerField.Kind())
//	controllerField.Elem().MethodByName(route.action)
//		method := appControllerPtr.MethodByName(route.action)
//		if !method.IsValid() {
//				LOGGER.Printf("E: Function %s not found on Controller %s",
//						route.action, route.controller)
//				http.NotFound(w, r)
//				return
//			}
//
//		resultValue := method.Call([]reflect.Value{ })[0]

//		result := resultValue.Interface().(*Response)
//		w.Write([]byte(result.content))

}
