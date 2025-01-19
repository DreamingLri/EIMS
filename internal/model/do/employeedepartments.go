// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Employeedepartments is the golang structure of table employeedepartments for DAO operations like Where/Data.
type Employeedepartments struct {
	g.Meta      `orm:"table:employeedepartments, do:true"`
	EdId        interface{} //
	EmpNo       interface{} //
	DeptNo      interface{} //
	EdEntrydate *gtime.Time //
	EdLeavedate *gtime.Time //
	EdStatus    interface{} //
}
