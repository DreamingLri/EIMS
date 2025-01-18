package v1

import (
	"demo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/department" method:"GET" tags:"Department"`
}

type GetListRes struct {
	List []entity.Departments `json:"list"`
}

type AddReq struct {
	g.Meta          `path:"/department" method:"POST" tags:"Department"`
	DeptName        string `json:"deptName"        orm:"dept_name"        description:""` //
	DeptPeoplecount int    `json:"deptPeoplecount" orm:"dept_peoplecount" description:""` //
}

type AddRes struct {
	DeptNo int `json:"deptNo"          orm:"dept_no"          description:""` //
}

type UpdateReq struct {
	g.Meta          `path:"/department" method:"PUT" tags:"Department"`
	DeptNo          int    `json:"deptNo"          orm:"dept_no"          description:""` //
	DeptName        string `json:"deptName"        orm:"dept_name"        description:""` //
	DeptPeoplecount int    `json:"deptPeoplecount" orm:"dept_peoplecount" description:""` //
}

type UpdateRes struct {
	DeptNo int `json:"deptNo"          orm:"dept_no"          description:""`
}

type DeleteReq struct {
	g.Meta `path:"/department" method:"DELETE" tags:"Department"`
	DeptNo int `json:"deptNo"          orm:"dept_no"          description:""`
}

type DeleteRes struct {
	DeptNo int `json:"deptNo"          orm:"dept_no"          description:""`
}
