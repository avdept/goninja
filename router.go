package goninja

import (
//	"fmt"
	"net/http"
	"reflect"
	"strings"
	"regexp"
)

type Params map[string]string

type Router struct {

	//slice with all routes
	routes []Route
	test string
	//hash with named route that matchers Route object
	named_routes map[string]*Route
}

var pattern string = "[a-zA-Z0-9]+"

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
	route := Route{method, pattern, action, controller, }
	CreateControllers(controller, c)
	r.routes = append(r.routes, route)
	return r
}

func (r *Router) RootRoute(action string, controller string, c interface{}) *Router {
	return r.Route("GET", "/", action, controller, c)
}


func (router *Router) match(request *http.Request) Route {

	var route Route

	if strings.Contains(request.URL.Path, "/assets/") {
		return route
	}

	RouteLoop:
	for _, r := range router.routes  {

		LOGGER.Println(request.URL.Path)

		if request.URL.Path == "/"{
			if r.pattern == "/" {
				route = r//root url
				break RouteLoop
			} else {
				continue RouteLoop
			}
		}

		route_pieces:= strings.Split(r.pattern, "/")
		request_pieces := strings.Split(request.URL.Path, "/")

		for i, x:= range route_pieces {
			if len(x) > 0 && string(x[0]) == ":" {
				route_pieces[i] = pattern
			}

		}
		if len(route_pieces) != len(request_pieces) {
			continue RouteLoop
		}
		route_regexp_string := strings.Join(route_pieces, "/")
		if route_regexp_string != "/" && regexp.MustCompile(route_regexp_string).MatchString(request.URL.Path) {  //WE have matched controller
			route = r
			break RouteLoop
		}
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

func isAssetRequest(path string) bool  {
	res := false
	//TODO add configurator
	if strings.Contains(path, "/assets/css/") || strings.Contains(path, "/assets/js/") || strings.Contains(path, "/assets/fonts/") {
		res = true
	}
	return res
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
			LOGGER.Println(result.Content)
		}
	} else if isAssetRequest(r.URL.Path) == true  {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		http.ServeFile(w, r, CURRENT_DIR +  r.URL.Path)
	} else {
		route.ControllerNotFound(w, r)
	}
}
