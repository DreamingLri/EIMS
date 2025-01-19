// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Employeedepartments is the golang structure for table employeedepartments.
type Employeedepartments struct {
	EdId        int         `json:"edId"        orm:"ed_id"        description:""` //
	EmpNo       int         `json:"empNo"       orm:"emp_no"       description:""` //
	DeptNo      int         `json:"deptNo"      orm:"dept_no"      description:""` //
	EdEntrydate *gtime.Time `json:"edEntrydate" orm:"Ed_entrydate" description:""` //
	EdLeavedate *gtime.Time `json:"edLeavedate" orm:"Ed_leavedate" description:""` //
	EdStatus    int         `json:"edStatus"    orm:"Ed_Status"    description:""` //
}
