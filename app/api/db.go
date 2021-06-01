package api

import (
	"gf-app-demo/app/dao"
	"gf-app-demo/app/model"
	"gf-app-demo/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/util/guid"
)

var Db = dbApi{}

type dbApi struct{}

func (*dbApi) GetAll(r *ghttp.Request) {

	list, err := dao.Test.Get()
	if err != nil {
		response.JsonExit(r, 400, err.Error())
	}

	response.JsonExit(r, 200, "操作成功", list)
}

func (*dbApi) Insert(r *ghttp.Request) {

	m := &model.Test{
		Name:  guid.S(),
		Sex:   grand.N(1, 2),
		Birth: gtime.Now(),
	}

	res, err := dao.Test.Data(m).Insert()
	if err != nil {
		return
	}

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

func (*dbApi) Up(r *ghttp.Request) {

	res, err := dao.Test.Data(g.Map{
		dao.Test.Columns.Name: guid.S(),
		dao.Test.Columns.Sex:  grand.N(1, 2),
	}).Insert()

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

func (*dbApi) Del(r *ghttp.Request) {
	id := r.GetInt("id")
	result, err := dao.Test.Del(id)
	if err != nil {
		response.JsonExit(r, 200, err.Error(), result)
	}
	rs, err := result.RowsAffected()
	if err != nil {
		response.JsonExit(r, 200, err.Error(), result)
	}
	response.JsonExit(r, 200, "操作成功", rs)
}
