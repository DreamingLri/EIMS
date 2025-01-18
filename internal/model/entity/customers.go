// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Customers is the golang structure for table customers.
type Customers struct {
	CustomerID   int    `json:"customerID"   orm:"CustomerID"   description:""` //
	CustomerName string `json:"customerName" orm:"CustomerName" description:""` //
	Company      string `json:"company"      orm:"Company"      description:""` //
	Sex          string `json:"sex"          orm:"Sex"          description:""` //
	Age          int    `json:"age"          orm:"Age"          description:""` //
	Telephone    string `json:"telephone"    orm:"Telephone"    description:""` //
	Address      string `json:"address"      orm:"Address"      description:""` //
}
