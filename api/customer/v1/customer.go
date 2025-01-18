package v1

import (
	"demo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/customer" method:"GET" tags:"Customer"`
}
type GetListRes struct {
	List []entity.Customers `json:"list"`
}

type AddReq struct {
	g.Meta       `path:"/customer" method:"POST" tags:"Customer"`
	CustomerName string `json:"customerName" orm:"CustomerName" description:""` //
	Company      string `json:"company"      orm:"Company"      description:""` //
	Sex          string `json:"sex"          orm:"Sex"          description:""` //
	Age          int    `json:"age"          orm:"Age"          description:""` //
	Telephone    string `json:"telephone"    orm:"Telephone"    description:""` //
	Address      string `json:"address"      orm:"Address"      description:""` //
}

type AddRes struct {
	CustomerID int `json:"/customerID"   orm:"CustomerID"   description:""` //
}

type UpdateReq struct {
	g.Meta       `path:"/customer" method:"PUT" tags:"Customer"`
	CustomerID   int    `json:"customerID"   orm:"CustomerID"   description:""` //
	CustomerName string `json:"customerName" orm:"CustomerName" description:""` //
	Company      string `json:"company"      orm:"Company"      description:""` //
	Sex          string `json:"sex"          orm:"Sex"          description:""` //
	Age          int    `json:"age"          orm:"Age"          description:""` //
	Telephone    string `json:"telephone"    orm:"Telephone"    description:""` //
	Address      string `json:"address"      orm:"Address"      description:""` //
}

type UpdateRes struct {
	CustomerID int `json:"customerID"   orm:"CustomerID"   description:""`
}

type DeleteReq struct {
	g.Meta     `path:"/customer" method:"DELETE" tags:"Customer"`
	CustomerID int `json:"customerID"   orm:"CustomerID"   description:""`
}

type DeleteRes struct {
	CustomerID int `json:"customerID"   orm:"CustomerID"   description:""`
}
