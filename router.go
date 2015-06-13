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
	test string
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

func (r *Router) Route(method string, pattern string, action string, controller string, c interface{}) *Router {
	route := Route{method, pattern, action, controller}
	CreateControllers(controller, c)
	r.routes = append(r.routes, route)
	return r
}

func (r *Router) RootRoute(action string, controller string, c interface{}) *Router {
	return r.Route("GET", "/", action, controller, c)
}


func (router *Router) match(request *http.Request) Route {

	var route Route
	var wildRoute Route

	for _, r := range router.routes  {

		if request.URL.Path == "/" && r.pattern == "/"{
			return  r
		}

		if r.pattern == request.URL.Path{
			route = r
		}

		//Support for wildcard route. Returns only if no route matcher to current request
		if r.pattern == "*" {
			wildRoute = r
		}
	}

	if reflect.DeepEqual(route, Route{}) {
		return wildRoute
	}
	return route
}

func (route *Route) ControllerNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("Controller that matches \"" + r.URL.Path + "\" could not be found"))
}

func (route *Route) CheckRequestMethod(w http.ResponseWriter, r *http.Request) bool {
	err := false

	if route.method == "" {
		return err
	}

	if route.method != r.Method {
		err = true
		w.WriteHeader(405)
		w.Write([]byte("Unexpected request method: " + r.Method + ". Expected: " + route.method))
	}
	return err
}

func  (router *Router) Handle(w http.ResponseWriter, r *http.Request) {
	route := router.match(r)

	// TODO this might be need rethinked
	if route.CheckRequestMethod(w, r) {
		return
	}
	obj, ok:= LaunchController(route.controller)
	if ok {
		v := reflect.ValueOf(obj)
//		ctrl_type := reflect.TypeOf(obj)
		ctrl_field := v.Elem().Field(0)
		var C *Controller = &Controller{Request: r,	Writer: w,	Name: route.controller, Action: route.action}
		ctrl_field.Set(reflect.ValueOf(C))
		action := v.MethodByName(route.action)
		if !action.IsValid(){
			w.WriteHeader(404)
			w.Write([]byte("Action with name " + route.action+  " wasn't found in controller " + route.controller))
		} else {
			res := action.Call([]reflect.Value{ })[0]
			result := res.Interface().(Response)
			w.Write([]byte(result.Content))
		}
	} else {
		route.ControllerNotFound(w, r)
	}
}
