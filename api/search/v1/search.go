package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Employee struct {
	EmpNo     int         `json:"empNo"     orm:"emp_no"     description:""`             //
	FirstName string      `json:"firstName" orm:"first_name" description:""`             //
	LastName  string      `json:"lastName"  orm:"last_name"  description:""`             //
	Gender    int         `json:"gender"    orm:"gender"     description:""`             //
	HireDate  *gtime.Time `json:"hireDate"  orm:"hire_date"  description:""`             //
	Birthday  *gtime.Time `json:"birthday"  orm:"birthday"   description:""`             //
	Address   string      `json:"address"   orm:"address"    description:""`             //
	Telephone string      `json:"telephone" orm:"telephone"  description:""`             //
	DeptName  string      `json:"deptName"        orm:"dept_name"        description:""` //
}

type SearchReq struct {
	g.Meta  `path:"/search" method:"POST" tags:"Search"`
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type SearchRes struct {
	List []Employee `json:"list"`
}
