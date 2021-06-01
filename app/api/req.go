package api

import (
	"github.com/gogf/gf/net/ghttp"
)

var Req = reqApi{}

type reqApi struct{}

func (*reqApi) Get(r *ghttp.Request) {
	r.Response.Writeln(r.Get("name"))
}

//  GET POST
func (*reqApi) Post(r *ghttp.Request) {
	r.Response.Writeln(r.Get("name"))
}

func (*reqApi) Put(r *ghttp.Request) {
	r.Response.Writeln(r.Get("name"))
}
func (*reqApi) Delete(r *ghttp.Request) {
	r.Response.Writeln(r.Get("name"))
}
