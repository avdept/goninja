package goninja

import (
	"net/http"
	"sync"
	//	"fmt"
	"log"
	"os"
)

var CURRENT_DIR, err = os.Getwd()
var mutex sync.Mutex

var LOGGER = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func Run(r *Router) {

	http.HandleFunc("/", r.Handle)
	http.ListenAndServe(":3000", nil)
}

func InitRouter() *Router {
	r := NewRouter()
	return r
}
