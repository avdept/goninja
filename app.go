package goninja

import (
	"net/http"
	"fmt"
)


func Run(r *Router) {
	// router = &Router{}

	http.HandleFunc("/", r.handle)
	http.ListenAndServe(":3000", nil)
}

func InitRouter() *Router {
	r := NewRouter()
	return r
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.RequestURI)
}
