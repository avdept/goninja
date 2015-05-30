package goninja

import (
	"net/http"
	"fmt"
)


func Run() {
	router := Router{}
	router.New()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)



}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println(w)
	fmt.Println(r)
}
