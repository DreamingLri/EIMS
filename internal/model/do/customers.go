// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Customers is the golang structure of table customers for DAO operations like Where/Data.
type Customers struct {
	g.Meta       `orm:"table:customers, do:true"`
	CustomerID   interface{} //
	CustomerName interface{} //
	Company      interface{} //
	Sex          interface{} //
	Age          interface{} //
	Telephone    interface{} //
	Address      interface{} //
}
