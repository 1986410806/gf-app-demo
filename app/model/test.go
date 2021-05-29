package model

type Test struct {
	Id   int    `orm:"id"`
	Name string `orm:"name"`
	Sex  int    `orm:"sex"`
}
