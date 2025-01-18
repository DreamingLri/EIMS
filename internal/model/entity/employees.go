// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Employees is the golang structure for table employees.
type Employees struct {
	EmpNo     int         `json:"empNo"     orm:"emp_no"     description:""` //
	FirstName string      `json:"firstName" orm:"first_name" description:""` //
	LastName  string      `json:"lastName"  orm:"last_name"  description:""` //
	Gender    int         `json:"gender"    orm:"gender"     description:""` //
	HireDate  *gtime.Time `json:"hireDate"  orm:"hire_date"  description:""` //
	Birthday  *gtime.Time `json:"birthday"  orm:"birthday"   description:""` //
	Address   string      `json:"address"   orm:"address"    description:""` //
	Telephone string      `json:"telephone" orm:"telephone"  description:""` //
}
