package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/anistark/gorouter"
)

type user struct {
	name string
}

func main() {
	r := gorouter.New(fallThrough)                         // Our fallthrough route.
	r.Use(fooMiddleware, barMiddleware, gorouter.Static()) // add global/router level middleware to run on every route.
	r.GET("/", root)
	r.GET("/users", users, authMiddleware)
	r.GET("/users/edit", root)
	r.EnableLogging(os.Stdout)
	r.Run(":8080")
}

// Notice the Middleware has a return type. True means go to the next middleware. False
// means to stop right here. If you return false to end the request-response cycle you MUST
// write something back to the client, otherwise it will be left hanging.
func fooMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	fmt.Println("Foo!")
	return true
}

func barMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	fmt.Println("Bar!")
	return true
}

func authMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	// fmt.Println("Doing Auth here")
	u := user{name: r.URL.Query().Get("name")}
	fmt.Printf("%x\n", &u.name)
	gorouter.Set(r, "user", u)
	return true
}

func fallThrough(w http.ResponseWriter, r *http.Request, params url.Values) {
	http.Error(w, "For your safety, do not push it.", http.StatusNotFound)
}

func test(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Println(params)
	fmt.Fprintf(w, "Hello World!")
}

func root(w http.ResponseWriter, r *http.Request, params url.Values) {
	w.WriteHeader(200)
	w.Write([]byte("Root!"))
}

func users(w http.ResponseWriter, r *http.Request, params url.Values) {
	u := gorouter.Get(r, "user").(user)
	fmt.Printf("%x\n", &u.name)
	fmt.Println("user is: ", u.name)
	fmt.Fprint(w, "user is: ", u.name)
}

func userShow(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "Hi %s", params["name"]) // Notice we are able to get the username from the url resource. Quite handy!
}
