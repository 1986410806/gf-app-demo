package api

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"master_home_api/app/resources"
)

var Hello = helloApi{}

type helloApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Word(r *ghttp.Request) {
	_ = r.Response.WriteJson(resources.ResultData{
		Code: resources.SuccessCode,
		Msg:  "操作成功",
		Data: []int{1, 2, 3, 4, 5, 6},
	})
}

func (*helloApi) Db(r *ghttp.Request) {
	db := g.DB()

	//test_model := model.Test{
	//	Name: "name",
	//	Sex:  1,
	//}
	l, err := db.GetAll("select * from test;")
	// res, err := db.Table(test_model).Data(test_model).InsertIgnore()

	fmt.Printf("%#v", l, err)

	r.Response.WriteJson(l)
}
