# Goninja
Simple web View-Controller framework using Golang. At the current moment its in very early development. In future basic model layer will be added, so it'd be possible to import 3rd party orm. This framework mostly inspired by Rails, so method\function names might be same.

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

`c.Redirect(absolute_url String)` redirects to provided url with method `GET` and status code `301`. In plans to extend, allowing use user defined status codes. Another goal is create absolute urls from relative, hiding all logic into framework's code, and possibly usage of named routes




## Roadmap

1. Add more flexibility to `views` module. Make it respond with not just `html` but at least `json`, `xml`
2. Sessions
3. Create skeleton for application, and possibly command line tool, to generate project.
4. Move some data to config files.
5. Add more logging.
6. Check for orm's and embed some into project, adding model layer

    

