package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"

	"gopkg.in/mgo.v2/bson"
)

type TaskDAL struct {
	mongo *common.MongoSessionStruct
}

func (t *TaskDAL) GetAllTaskHeaders() (taskGetList map[string][]*types.TaskHeader_Get, err error) {
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
	taskGetList = make(map[string][]*types.TaskHeader_Get)

	var taskList []*types.TaskHeader
	t.mongo.Collection.Find(nil).All(&taskList)
	taskCount := len(taskList)
	taskGetList["data"] = make([]*types.TaskHeader_Get, taskCount, taskCount)
	for index, value := range taskList {
		taskGet := new(types.TaskHeader_Get)
		common.StructDeepCopy(value, taskGet)
		// 获取人员姓名
		emp := new(types.EmployeeName)
		err1 := t.mongo.Db.C("M_Employees").FindId(*value.PrimaryExecutorObjectID).One(&emp)
		if err1 == nil {
			taskGet.PrimaryExecutor = emp.Name
		}
		taskGetList["data"][index] = taskGet
	}
	return
}

func (t *TaskDAL) GetTaskDetail(id string) (taskGet map[string]*types.Task_Get, err error) {
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
	taskGet = make(map[string]*types.Task_Get)
	taskGet["data"] = new(types.Task_Get)
	task := new(types.Task)
	t.mongo.Collection.Find(bson.M{"id": id}).One(task)
	common.StructDeepCopy(task, taskGet["data"])

	// 获取人员姓名
	emp := new(types.EmployeeName)
	err1 := t.mongo.Db.C("M_Employees").FindId(task.PrimaryExecutorObjectID).One(&emp)
	if err1 == nil {
		taskGet["data"].PrimaryExecutor = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = t.mongo.Db.C("M_Employees").FindId(task.CreatorObjectID).One(&emp)
	if err1 == nil {
		taskGet["data"].Creator = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = t.mongo.Db.C("M_Employees").FindId(task.PrimarySellerObjectID).One(&emp)
	if err1 == nil {
		taskGet["data"].PrimarySeller = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = t.mongo.Db.C("M_Employees").FindId(task.PrimaryOCObjectID).One(&emp)
	if err1 == nil {
		taskGet["data"].PrimaryOC = emp.Name
	}
	otherExecutorsCount := len(task.OtherExecutorObjectIds)
	taskGet["data"].OtherExecutors = make([]string, otherExecutorsCount, otherExecutorsCount)
	for index, value := range task.OtherExecutorObjectIds {
		emp = new(types.EmployeeName)
		err1 = t.mongo.Db.C("M_Employees").FindId(value).One(&emp)
		if err1 == nil {
			taskGet["data"].OtherExecutors[index] = *emp.Name
		}
	}

	product := new(types.ProductName)
	err1 = t.mongo.Db.C("T_Products").FindId(task.ParentProductObjectID).One(&emp)
	if err1 == nil {
		taskGet["data"].ParentProduct = product.Name
	}
	project := new(types.ProjectName)
	err1 = t.mongo.Db.C("T_Projects").FindId(task.ParentProjectObjectID).One(&emp)
	if err1 == nil {
		taskGet["data"].ParentProject = project.Name
	}
	return
}

func (t *TaskDAL) AddTask(taskPost types.Task_Post) (s map[string]map[string]string, err error) {
	t.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer t.mongo.CloseSession()

	task := new(types.Task)
	common.StructDeepCopy(taskPost, task)
	task.OID = bson.NewObjectId()
	t.mongo.UseDB("local")

	err = t.mongo.UseCollection("T_Tasks")
	if err != nil {
		return
	}

	err = t.mongo.Collection.Insert(task)
	if err != nil {
		return
	}
	s = make(map[string]map[string]string)
	s["data"] = make(map[string]string)
	s["data"]["Id"] = *task.ID

	return
}
