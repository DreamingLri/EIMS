// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo/internal/dao/internal"
)

// internalEmployeedepartmentsDao is internal type for wrapping internal DAO implements.
type internalEmployeedepartmentsDao = *internal.EmployeedepartmentsDao

// employeedepartmentsDao is the data access object for table employeedepartments.
// You can define custom methods on it to extend its functionality as you wish.
type employeedepartmentsDao struct {
	internalEmployeedepartmentsDao
}

var (
	// Employeedepartments is globally public accessible object for table employeedepartments operations.
	Employeedepartments = employeedepartmentsDao{
		internal.NewEmployeedepartmentsDao(),
	}
)

// Fill with you ideas below.
