package goninja

import("net/http"
	"time"
	)

//TODO Create universal way to create sessions for any kind of resource(user, etc) using its id as value appending its kind(User, Administrator). Also encode it. Create method that would return back its resource

func CreateSession(actor interface{}, request http.Request, writer http.ResponseWriter) {
	c := http.Cookie{"test", "tcookie", "/", request.Host, time.Now().Add(10), "20000", 86400, false, true, "test=tcookie", []string{"test=tcookie"}}
	http.SetCookie(writer, &c)
}
