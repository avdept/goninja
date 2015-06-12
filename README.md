# goninja
Simple web framework using Golang

## Routes

In App skeleton you will find file `routes.go` which is responsible for routes registering. Args that required to register it:

`"GET"` - Request method

`"/"` - Request url

`"Index"` - Action that will respond

`"AppController"` - Controller name

`&controller.AppController` - Pointer to controller struct. Any fields can be passed, but only exported type. Can be passed as empty struct aswell

    func Register_routes(r *goninja.Router){
      r.AddRoute("GET", "/", "Index", "AppController", &controller.AppController{Name: "Hello"})    //example of registering router
    }    
