package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetListReq struct {
	g.Meta `path:"/management" method:"GET" tags:"Management"`
}

type GetListRes struct {
	EmpNo       int         `json:"empNo"       orm:"emp_no"       description:""` //
	DeptNo      int         `json:"deptNo"      orm:"dept_no"      description:""` //
	EdEntrydate *gtime.Time `json:"edEntrydate" orm:"Ed_entrydate" description:""` //
	EdLeavedate *gtime.Time `json:"edLeavedate" orm:"Ed_leavedate" description:""` //
	EdStatus    int         `json:"edStatus"    orm:"Ed_Status"    description:""` //
}

type UpdateReq struct {
	g.Meta      `path:"/management" method:"PUT" tags:"Management"`
	EdId        int         `json:"edId"        orm:"ed_id"        description:""` //
	EmpNo       int         `json:"empNo"       orm:"emp_no"       description:""` //
	DeptNo      int         `json:"deptNo"      orm:"dept_no"      description:""` //
	EdEntrydate *gtime.Time `json:"edEntrydate" orm:"Ed_entrydate" description:""` //
	EdLeavedate *gtime.Time `json:"edLeavedate" orm:"Ed_leavedate" description:""` //
	EdStatus    int         `json:"edStatus"    orm:"Ed_Status"    description:""` //
}

type UpdateRes struct {
	EdId int `json:"edId"        orm:"ed_id"        description:""`
}

type DeleteReq struct {
	g.Meta `path:"/management" method:"DELETE" tags:"Management"`
	EdId   int `json:"edId"        orm:"ed_id"        description:""`
}

type DeleteRes struct {
	EdId int `json:"edId"        orm:"ed_id"        description:""`
}
