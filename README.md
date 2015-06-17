# Goninja
Simple web View-Controller framework using Golang. At the current moment its in very early development. In future basic model layer will be added, so it'd be possible to import 3rd party orm

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


## Controllers

All controllers should be inside folder `app/controllers`. At the moment there is no difference for naming. Probably will be added in future.

This is example of controller file

    package controller

    import "github.com/avdept/goninja"

    type AppController struct {
      *goninja.Controller
    }

    func (c AppController) Index() goninja.Response{
      return c.Render()
    }
    
    
The controller function should always have `goninja.Response` as return type.This is needed for type assertion inside router. In future might be improved.

`c.Render()` returns `goninja.Response` by itself. Also this function renders `html` template from `app/views/controller_name/action_name.html`. `controller_name` and `action_name` are taken from `routes` file.






## Roadmap

1. Add more flexibility to `views` module. Make it respond with not just `html` but at least `json`, `xml`
2. Create router, that would match urls not just by name, but by mask(e.g. passing `id` - `users_controller/:id/edit`)
3. Sessions
4. Create skeleton for application, and possibly command line tool, to generate project.
5. Move some data to config files.
6. Add more logging.
7. Check for orm's and embed some into project, adding model layer

    

