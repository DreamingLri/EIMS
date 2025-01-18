package model

import (
	"demo/internal/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetEmployeeListInput struct{}

type GetEmployeeListOutput struct {
	List []entity.Employees `json:"customer_list"`
}

type AddEmployeeInput struct {
	FirstName string      `json:"firstName" orm:"first_name" description:""` //
	LastName  string      `json:"lastName"  orm:"last_name"  description:""` //
	Gender    int         `json:"gender"    orm:"gender"     description:""` //
	HireDate  *gtime.Time `json:"hireDate"  orm:"hire_date"  description:""` //
	Birthday  *gtime.Time `json:"birthday"  orm:"birthday"   description:""` //
	Address   string      `json:"address"   orm:"address"    description:""` //
	Telephone string      `json:"telephone" orm:"telephone"  description:""` //
}

type AddEmployeeOutput struct {
	EmpNo int `json:"empNo"     orm:"emp_no"     description:""` //
}

type UpdateEmployeeInput struct {
	EmpNo     int         `json:"empNo"     orm:"emp_no"     description:""` //
	FirstName string      `json:"firstName" orm:"first_name" description:""` //
	LastName  string      `json:"lastName"  orm:"last_name"  description:""` //
	Gender    int         `json:"gender"    orm:"gender"     description:""` //
	HireDate  *gtime.Time `json:"hireDate"  orm:"hire_date"  description:""` //
	Birthday  *gtime.Time `json:"birthday"  orm:"birthday"   description:""` //
	Address   string      `json:"address"   orm:"address"    description:""` //
	Telephone string      `json:"telephone" orm:"telephone"  description:""` //
}

type UpdateEmployeeOutput struct {
	EmpNo int `json:"empNo"     orm:"emp_no"     description:""` //
}

type DeleteEmployeeInput struct {
	EmpNo int `json:"empNo"     orm:"emp_no"     description:""` //
}

type DeleteEmployeeOutput struct{}
