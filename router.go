package goninja

import (
	"fmt"
)

type Router struct {

}

func (router *Router) New() {
	fmt.Println("Router initialized")
}
