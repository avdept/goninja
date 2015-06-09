package goninja

import (
	"net/http"
	"fmt"
	"log"
	"os"
)

var LOGGER = log.New(os.Stdout, "", log.Ldate | log.Ltime | log.Lshortfile)

func Run(r *Router) {
	// router = &Router{}

	http.HandleFunc("/", r.Handle)
	http.ListenAndServe(":3000", nil)
}

func InitRouter() *Router {
	r := NewRouter()
	return r
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.RequestURI)
}
