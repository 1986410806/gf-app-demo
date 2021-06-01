// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"database/sql"
	"gf-app-demo/app/dao/internal"
	"github.com/gogf/gf/database/gdb"
)

// testDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type testDao struct {
	internal.TestDao
}

var (
	// Test is globally public accessible object for table test operations.
	Test = testDao{
		internal.Test,
	}
)

// Fill with you ideas below.

// 创建
func (t *testDao) create(data interface{}) (res sql.Result, error error) {
	return t.Data(data).Insert()
}

// 更新
func (t *testDao) update(data interface{}) (res sql.Result, error error) {
	return t.Data(data).Update()
}

// Get 查询
func (t *testDao) Get() (res gdb.Result, error error) {
	return t.FindAll()
}

// 删除
func (t *testDao) Del(id int) (res sql.Result, error error) {
	return t.Delete("id=?", id)
}
