package v1

import (
	"demo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetListReq struct {
	g.Meta `path:"/employee" method:"GET" tags:"Employee"`
}

type GetListRes struct {
	List []entity.Employees `json:"list"`
}

type AddReq struct {
	FirstName string      `json:"firstName" orm:"first_name" description:""` //
	LastName  string      `json:"lastName"  orm:"last_name"  description:""` //
	Gender    int         `json:"gender"    orm:"gender"     description:""` //
	HireDate  *gtime.Time `json:"hireDate"  orm:"hire_date"  description:""` //
	Birthday  *gtime.Time `json:"birthday"  orm:"birthday"   description:""` //
	Address   string      `json:"address"   orm:"address"    description:""` //
	Telephone string      `json:"telephone" orm:"telephone"  description:""` //
}

type AddRes struct {
	EmpNo int `json:"empNo"     orm:"emp_no"     description:""` //
}

type UpdateReq struct {
	g.Meta    `path:"/employee" method:"PUT" tags:"Employee"`
	EmpNo     int         `json:"empNo"     orm:"emp_no"     description:""` //
	FirstName string      `json:"firstName" orm:"first_name" description:""` //
	LastName  string      `json:"lastName"  orm:"last_name"  description:""` //
	Gender    int         `json:"gender"    orm:"gender"     description:""` //
	HireDate  *gtime.Time `json:"hireDate"  orm:"hire_date"  description:""` //
	Birthday  *gtime.Time `json:"birthday"  orm:"birthday"   description:""` //
	Address   string      `json:"address"   orm:"address"    description:""` //
	Telephone string      `json:"telephone" orm:"telephone"  description:""` //
}

type UpdateRes struct {
	EmpNo int `json:"empNo"     orm:"emp_no"     description:""`
}

type DeleteReq struct {
	g.Meta `path:"/employee" method:"DELETE" tags:"Employee"`
	EmpNo  int `json:"empNo"     orm:"emp_no"     description:""` //
}

type DeleteRes struct {
	EmpNo int `json:"empNo"     orm:"emp_no"     description:""`
}
