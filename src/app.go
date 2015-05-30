package goninja

import (
	"net/http"
	"fmt"
)


func Run() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)


}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Initial start")
	fmt.Println(w)
	fmt.Println(r)
}
