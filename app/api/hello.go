package api

import (
	"gf-app-demo/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {
	r.Response.Writeln("Hello World! h")
}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Word(r *ghttp.Request) {
	response.JsonExit(r, 100, "hahahaa")
}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Exit(r *ghttp.Request) {
	go func() {
		r.Response.WriteExit("starting...")
		panic("exit")
	}()
	r.Response.WriteExit("success")
}
