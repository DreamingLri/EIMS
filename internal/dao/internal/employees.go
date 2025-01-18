// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmployeesDao is the data access object for table employees.
type EmployeesDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns EmployeesColumns // columns contains all the column names of Table for convenient usage.
}

// EmployeesColumns defines and stores column names for table employees.
type EmployeesColumns struct {
	EmpNo     string //
	FirstName string //
	LastName  string //
	Gender    string //
	HireDate  string //
	Birthday  string //
	Address   string //
	Telephone string //
}

// employeesColumns holds the columns for table employees.
var employeesColumns = EmployeesColumns{
	EmpNo:     "emp_no",
	FirstName: "first_name",
	LastName:  "last_name",
	Gender:    "gender",
	HireDate:  "hire_date",
	Birthday:  "birthday",
	Address:   "address",
	Telephone: "telephone",
}

// NewEmployeesDao creates and returns a new DAO object for table data access.
func NewEmployeesDao() *EmployeesDao {
	return &EmployeesDao{
		group:   "default",
		table:   "employees",
		columns: employeesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EmployeesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EmployeesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EmployeesDao) Columns() EmployeesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EmployeesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EmployeesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EmployeesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
