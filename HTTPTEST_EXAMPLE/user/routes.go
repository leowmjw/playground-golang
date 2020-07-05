package user

import (
	"github.com/dimfeld/httptreemux/v5"
)

type Server struct {
	router *Router
}

func (s *Server) routes() {
	// Add routes .. parameters passed in via contextParams
	s.router.HandleFunc("GET", "/api/users", handlerUserAuth())
	s.router.HandleFunc("GET", "/api/user/:userId", handlerUserAuth())
	// More patterns here:  https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html
}

type ComplexServer struct {
	router *httptreemux.TreeMux
}

func (c *ComplexServer) routes() {
	cg := c.router.NewGroup("/api")
	cg.UsingContext().GET("/users", handlerUserAuth())
	cg.UsingContext().GET("/user/:userId", handlerUserAuth())
}
