GoRouter
=======

## Installation `go get github.com/anistark/gorouter`


## Personalisation in progress


#########################################################################


## Usage

Initialising Router: `r := gorouter.New()`


* Declaring universal middleware for all routers: `r.Use(fooMiddleware, barMiddleware, gorouter.Static())`

* GET: `r.GET("/", root)`

* GET with Auth Middleware: `r.GET("/users", users, authMiddleware)`


[Initial Fork](https://github.com/acmacalister/helm)
