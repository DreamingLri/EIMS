package model

import "demo/internal/model/entity"

type GetCustomerListInput struct{}

type GetCustomerListOutput struct {
	List []entity.Customers `json:"customer_list"`
}

type AddCustomerInput struct {
	CustomerName string `json:"customerName" orm:"CustomerName" description:""` //
	Company      string `json:"company"      orm:"Company"      description:""` //
	Sex          string `json:"sex"          orm:"Sex"          description:""` //
	Age          int    `json:"age"          orm:"Age"          description:""` //
	Telephone    string `json:"telephone"    orm:"Telephone"    description:""` //
	Address      string `json:"address"      orm:"Address"      description:""` //
}

type AddCustomerOutput struct {
	CustomerID int `json:"customerID"   orm:"CustomerID"   description:""` //
}

type UpdateCustomerInput struct {
	CustomerID   int    `json:"customerID"   orm:"CustomerID"   description:""` //
	CustomerName string `json:"customerName" orm:"CustomerName" description:""`
	Company      string `json:"company"      orm:"Company"      description:""`
	Sex          string `json:"sex"          orm:"Sex"          description:""` //
	Age          int    `json:"age"          orm:"Age"          description:""` //
	Telephone    string `json:"telephone"    orm:"Telephone"    description:""` //
	Address      string `json:"address"      orm:"Address"      description:""` //
}

type UpdateCustomerOutput struct {
	CustomerID int `json:"customerID"   orm:"CustomerID"   description:""`
}

type DeleteCustomerInput struct {
	CustomerID int `json:"customerID"   orm:"CustomerID"   description:""`
}

type DeleteCustomerOutput struct{}
