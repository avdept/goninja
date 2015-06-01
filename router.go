package goninja

import (
	"fmt"
	"net/http"
	// "strings"
	// "html/template"
	// "regexp"
)

type Params map[string]string

type Handler interface {}


type Router struct {

	//slice with all routes
	routes []Route
	//hash with named route that matchers Route object
	named_routes map[string]*Route


}

type Route struct {

	method string

	pattern string

	handler Handler

}

func NewRouter() *Router {
	router := &Router{}
	return router
}

func (r *Router) addRoute(method string, pattern string, handler *Handler) *Router {
	route := Route{method: method, pattern: pattern, handler: handler}
	r.routes = append(r.routes, route)
	return r
}



func  (*Router) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("served from router")
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Println("Served root path")
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Hello", "world")
}
