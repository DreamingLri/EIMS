// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Departments is the golang structure of table departments for DAO operations like Where/Data.
type Departments struct {
	g.Meta          `orm:"table:departments, do:true"`
	DeptNo          interface{} //
	DeptName        interface{} //
	DeptPeoplecount interface{} //
}
