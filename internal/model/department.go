package model

import "demo/internal/model/entity"

type GetDepartmentListInput struct{}

type GetDepartmentListOutput struct {
	List []entity.Departments `json:"customer_list"`
}

type AddDepartmentInput struct {
	DeptName        string `json:"deptName"        orm:"dept_name"        description:""` //
	DeptPeoplecount int    `json:"deptPeoplecount" orm:"dept_peoplecount" description:""` //
}

type AddDepartmentOutput struct {
	DeptNo int `json:"deptNo"          orm:"dept_no"          description:""` //
}

type UpdateDepartmentInput struct {
	DeptNo          int    `json:"deptNo"          orm:"dept_no"          description:""` //
	DeptName        string `json:"deptName"        orm:"dept_name"        description:""` //
	DeptPeoplecount int    `json:"deptPeoplecount" orm:"dept_peoplecount" description:""` //
}

type UpdateDepartmentOutput struct {
	DeptNo int `json:"deptNo"          orm:"dept_no"          description:""` //
}

type DeleteDepartmentInput struct {
	DeptNo int `json:"deptNo"          orm:"dept_no"          description:""` //
}

type DeleteDepartmentOutput struct{}
