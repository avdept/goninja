package goninja

import (
//	"fmt"
	"net/http"
	"reflect"
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

func (r *Router) AddRoute(method string, pattern string, action string, controller string, c interface{}) *Router {
	route := Route{method, pattern, action, controller}
	CreateControllers(controller, c)
	r.routes = append(r.routes, route)
	return r
}

func (router *Router) match(r *http.Request) Route {

	//This is temporary stub matcher, to test handling of different actions
	return router.routes[1]
}

func  (router *Router) Handle(w http.ResponseWriter, r *http.Request) {
	route := router.match(r)
	obj := LaunchController(route.controller)
	v := reflect.ValueOf(obj)
	ctrl_type := reflect.TypeOf(obj)
	ctrl_field := v.Elem().FieldByName("Ctrl")

	var C *Controller = &Controller{Request: r,	Writer: w,	Name: ctrl_type.Name()	}

	ctrl_field.Set(reflect.ValueOf(C))
	action := v.MethodByName(route.action)
	if !action.IsValid(){
		w.WriteHeader(404)
		w.Write([]byte("Action with name " + route.action+  " wasn't found in controller " + route.controller))
	} else {
		res := action.Call([]reflect.Value{ })[0]
		result := res.Interface().(*Response)
		w.Write([]byte(result.Content))
	}
}
