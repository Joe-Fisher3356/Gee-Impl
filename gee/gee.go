package gee

import (
	"fmt"
	"net/http"
)

// Define a function type
type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	fmt.Println("NewEngine() called")
	return &Engine{router: newRouter()}
}

func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c:= newContext(&w, req)
	e.router.handle(c)
}
