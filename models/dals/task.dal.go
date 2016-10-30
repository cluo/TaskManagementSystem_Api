package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"
	"strconv"
	"strings"

	"github.com/gocircuit/circuit/use/errors"

	"time"

	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// TaskDAL 定义
type TaskDAL struct {
	mongo *common.MongoSessionStruct
}

// GetAllTaskHeaders 定义
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
		if value.PrimaryExecutorObjectID != nil {
			err1 := t.mongo.Db.C("M_Employees").FindId(*value.PrimaryExecutorObjectID).One(&emp)
			if err1 == nil {
				taskGet.PrimaryExecutor = emp.Name
			}
		}
		taskGetList["data"][index] = taskGet
	}
	return
}

// GetTaskDetail 定义
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

// AddTask 定义
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

	dateString := time.Now().Format("20060102")
	var maxID *types.MaxID
	taskRegex := "^T" + dateString + "[0-9]{4}$"
	iter := t.mongo.Collection.Find(bson.M{"id": bson.M{"$regex": taskRegex}}).Sort("-id").Iter()
	iter.Next(&maxID)
	var id string
	if maxID == nil {
		id = "T" + dateString + "0001"
	} else {
		maxNum, _ := strconv.Atoi((*maxID.ID)[9:])
		if maxNum >= 9999 {
			err = errors.NewError("今天新建的任务超出最大限度")
			return
		}
		id = fmt.Sprintf("T%s%04d", dateString, maxNum+1)
	}
	task.ID = &id

	err = t.mongo.Collection.Insert(task)
	if err != nil && strings.Contains(err.Error(), "E11000 duplicate key error collection:") {
		return t.AddTask(taskPost)
	} else if err != nil {
		return
	}

	s = make(map[string]map[string]string)
	s["data"] = make(map[string]string)
	s["data"]["Id"] = *task.ID

	return
}
