package blls

import (
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
)

type EmployeeBLL struct {
}

func (bll *EmployeeBLL) GetAllEmployees() (e []*types.Employee_Get, err error) {
	return (&dals.EmployeeDAL{}).GetAllEmployees()
}
