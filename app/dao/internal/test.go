// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// TestDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type TestDao struct {
	gmvc.M              // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB      // DB is the raw underlying database management object.
	Table   string      // Table is the table name of the DAO.
	Columns testColumns // Columns contains all the columns of Table that for convenient usage.
}

// TestColumns defines and stores column names for table test.
type testColumns struct {
	Id   string //
	Name string //
	Sex  string //
}

var (
	// Test is globally public accessible object for table test operations.
	Test = TestDao{
		M:     g.DB("default").Model("test").Safe(),
		DB:    g.DB("default"),
		Table: "test",
		Columns: testColumns{
			Id:   "id",
			Name: "name",
			Sex:  "sex",
		},
	}
)
