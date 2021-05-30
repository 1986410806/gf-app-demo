package api

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"master_home_api/app/dao"
	"master_home_api/app/model/test"
)

var Db = dbApi{}

type dbApi struct{}

func (*dbApi) GetAll(r *ghttp.Request) {
	db := g.DB()

	//test_model := model.Test{
	//	Name: "name",
	//	Sex:  1,
	//}
	l, err := db.GetAll("select * from test;")
	// res, err := db.Table(test_model).Data(test_model).InsertIgnore()

	fmt.Printf("%+v %s", l, err)

	r.Response.WriteJsonExit(l)
}

func (*dbApi) Insert(r *ghttp.Request) {

	testModel := test.Entity{
		Name: "2ts",
		Sex:  11,
	}

	res, err := dao.Test.Insert(testModel)

	if err != nil {
		r.Response.Write(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return
	}
	err = r.Response.WriteJson(id)
	if err != nil {
		r.Response.Write(err)
	}
}
