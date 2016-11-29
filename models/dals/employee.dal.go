package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"
)

// EmployeeDAL 定义
type EmployeeDAL struct {
	mongo *common.MongoSessionStruct
}

// GetAllEmployees 定义
func (dal *EmployeeDAL) GetAllEmployees() (e []*types.Employee_Get, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")

	err = dal.mongo.UseCollection("M_Employees")
	if err != nil {
		return
	}

	var employeeList []*types.EmployeeInfo
	err = dal.mongo.Collection.Find(nil).Sort("name").All(&employeeList)
	if err != nil {
		return
	}
	employeeCount := len(employeeList)
	e = make([]*types.Employee_Get, employeeCount, employeeCount)
	for index, value := range employeeList {
		employeeGet := new(types.Employee_Get)
		common.StructDeepCopy(value, employeeGet)
		dept := new(types.DeptName)
		err1 := dal.mongo.Db.C("M_Departments").FindId(value.DeptObjectID).One(&dept)
		if err1 == nil {
			employeeGet.Dept = dept.Name
		}
		e[index] = employeeGet
	}
	return
}
