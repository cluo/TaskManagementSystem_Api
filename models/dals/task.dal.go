package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"

	"gopkg.in/mgo.v2/bson"
)

type TaskDAL struct {
	mongo *common.MongoSessionStruct
}

func (t *TaskDAL) GetAllTaskHeaders() (taskList map[string][]*types.TaskHeader, err error) {
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
	taskList = make(map[string][]*types.TaskHeader)
	task := new(types.TaskHeader)
	taskList["data"] = make([]*types.TaskHeader, 0, 10)

	iter := t.mongo.Collection.Find(nil).Iter()
	for iter.Next(&task) {
		// taskList[task.ID] = task
		taskList["data"] = append(taskList["data"], task)
		task = new(types.TaskHeader)
	}

	// 获取人员姓名
	for _, value := range taskList["data"] {
		emp := new(types.EmployeeName)
		err1 := t.mongo.Db.C("M_Employees").FindId(value.PrimaryExecutorObjectID).One(&emp)
		if err1 == nil {
			value.PrimaryExecutor = emp.Name
		}
	}
	return
}
func (t *TaskDAL) GetTaskDetail(id string) (task map[string]*types.Task, err error) {
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

	task = make(map[string]*types.Task)
	task["data"] = new(types.Task)
	t.mongo.Collection.Find(bson.M{"id": id}).One(task["data"])

	// 获取人员姓名
	emp := new(types.EmployeeName)
	err1 := t.mongo.Db.C("M_Employees").FindId(task["data"].PrimaryExecutorObjectID).One(&emp)
	if err1 == nil {
		task["data"].PrimaryExecutor = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = t.mongo.Db.C("M_Employees").FindId(task["data"].CreatorObjectID).One(&emp)
	if err1 == nil {
		task["data"].Creator = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = t.mongo.Db.C("M_Employees").FindId(task["data"].PrimarySellerObjectID).One(&emp)
	if err1 == nil {
		task["data"].PrimarySeller = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = t.mongo.Db.C("M_Employees").FindId(task["data"].PrimaryOCObjectID).One(&emp)
	if err1 == nil {
		task["data"].PrimaryOC = emp.Name
	}
	otherExecutorsCount := len(task["data"].OtherExecutorObjectIds)
	task["data"].OtherExecutors = make([]string, otherExecutorsCount, otherExecutorsCount)
	for index, value := range task["data"].OtherExecutorObjectIds {
		emp = new(types.EmployeeName)
		err1 = t.mongo.Db.C("M_Employees").FindId(value).One(&emp)
		if err1 == nil {
			task["data"].OtherExecutors[index] = emp.Name
		}
	}

	product := new(types.ProductName)
	err1 = t.mongo.Db.C("T_Products").FindId(task["data"].ParentProductObjectID).One(&emp)
	if err1 == nil {
		task["data"].ParentProduct = product.Name
	}
	project := new(types.ProjectName)
	err1 = t.mongo.Db.C("T_Projects").FindId(task["data"].ParentProjectObjectID).One(&emp)
	if err1 == nil {
		task["data"].ParentProject = project.Name
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
