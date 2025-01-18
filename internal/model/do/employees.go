// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Employees is the golang structure of table employees for DAO operations like Where/Data.
type Employees struct {
	g.Meta    `orm:"table:employees, do:true"`
	EmpNo     interface{} //
	FirstName interface{} //
	LastName  interface{} //
	Gender    interface{} //
	HireDate  *gtime.Time //
	Birthday  *gtime.Time //
	Address   interface{} //
	Telephone interface{} //
}
