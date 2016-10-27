package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"

	"gopkg.in/mgo.v2/bson"
)

type TaskDAL struct {
	mongo *common.MongoSessionStruct
}

func (t *TaskDAL) GetAllTaskHeaders() (taskList map[string]*types.TaskHeader, err error) {
	t.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer t.mongo.CloseSession()
	t.mongo.UseDB("local")
	err = t.mongo.UseCollection("T_Tasks")
	if err != nil {
		return
	}

	task := new(types.TaskHeader)
	taskList = make(map[string]*types.TaskHeader)

	iter := t.mongo.Collection.Find(nil).Iter()
	for iter.Next(&task) {
		taskList[task.ID] = task
		task = new(types.TaskHeader)
	}

	// 获取人员姓名
	for _, value := range taskList {
		emp := new(types.EmployeeName)
		err1 := t.mongo.Db.C("M_Employees").FindId(value.PrimaryExecutorObjectID).One(&emp)
		if err1 == nil {
			value.PrimaryExecutor = emp.Name
		}
	}
	return
}
func (t *TaskDAL) GetTaskDetail(id string) (task types.Task, err error) {
	t.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer t.mongo.CloseSession()
	t.mongo.UseDB("local")
	err = t.mongo.UseCollection("T_Tasks")
	if err != nil {
		return
	}

	t.mongo.Collection.Find(bson.M{"id": id}).One(&task)

	// 获取人员姓名
	emp := new(types.EmployeeName)
	err1 := t.mongo.Db.C("M_Employees").FindId(task.PrimaryExecutorObjectID).One(&emp)
	if err1 == nil {
		task.PrimaryExecutor = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = t.mongo.Db.C("M_Employees").FindId(task.CreatorObjectID).One(&emp)
	if err1 == nil {
		task.Creator = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = t.mongo.Db.C("M_Employees").FindId(task.PrimarySellerObjectID).One(&emp)
	if err1 == nil {
		task.PrimarySeller = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = t.mongo.Db.C("M_Employees").FindId(task.PrimaryOCObjectID).One(&emp)
	if err1 == nil {
		task.PrimaryOC = emp.Name
	}
	otherExecutorsCount := len(task.OtherExecutorObjectIds)
	task.OtherExecutors = make([]string, otherExecutorsCount, otherExecutorsCount)
	for index, value := range task.OtherExecutorObjectIds {
		emp = new(types.EmployeeName)
		err1 = t.mongo.Db.C("M_Employees").FindId(value).One(&emp)
		if err1 == nil {
			task.OtherExecutors[index] = emp.Name
		}
	}

	product := new(types.ProductName)
	err1 = t.mongo.Db.C("T_Products").FindId(task.ParentProductObjectID).One(&emp)
	if err1 == nil {
		task.ParentProduct = product.Name
	}
	project := new(types.ProjectName)
	err1 = t.mongo.Db.C("T_Projects").FindId(task.ParentProjectObjectID).One(&emp)
	if err1 == nil {
		task.ParentProject = project.Name
	}
	return
}

// {
//   "otherExecutorObjectIds": [
//     "580dec5bacfecf773c1ec328",
//     "580dfc4eacfecf773c1ec33a"
//   ],
//   "otherExecutors": null,
//   "parentProductObjectId": "",
//   "parentProduct": "",
//   "parentProjectObjectId": "",
//   "parentProject": ""
// }
