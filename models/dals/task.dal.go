package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// TaskDAL 定义
type TaskDAL struct {
	mongo *common.MongoSessionStruct
}

// GetAllTaskHeaders 定义
func (dal *TaskDAL) GetAllTaskHeaders() (taskGetList []*types.TaskHeader_Get, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Tasks")
	if err != nil {
		return
	}

	var taskList []*types.TaskHeader
	dal.mongo.Collection.Find(nil).All(&taskList)
	taskCount := len(taskList)
	taskGetList = make([]*types.TaskHeader_Get, taskCount, taskCount)
	for index, value := range taskList {
		taskGet := new(types.TaskHeader_Get)
		common.StructDeepCopy(value, taskGet)
		// 获取人员姓名
		emp := new(types.EmployeeName)
		if value.PrimaryExecutorObjectID != nil {
			err1 := dal.mongo.Db.C("M_Employees").FindId(*value.PrimaryExecutorObjectID).One(&emp)
			if err1 == nil {
				taskGet.PrimaryExecutor = emp.Name
			}
		}
		taskGetList[index] = taskGet
	}
	return
}

// GetTaskDetail 定义
func (dal *TaskDAL) GetTaskDetail(id string) (taskGet *types.Task_Get, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Tasks")
	if err != nil {
		return
	}
	taskGet = new(types.Task_Get)
	task := new(types.Task)
	dal.mongo.Collection.Find(bson.M{"id": id}).One(task)
	common.StructDeepCopy(task, taskGet)

	// 获取人员姓名
	emp := new(types.EmployeeName)
	err1 := dal.mongo.Db.C("M_Employees").FindId(task.PrimaryExecutorObjectID).One(&emp)
	if err1 == nil {
		taskGet.PrimaryExecutor = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = dal.mongo.Db.C("M_Employees").FindId(task.CreatorObjectID).One(&emp)
	if err1 == nil {
		taskGet.Creator = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = dal.mongo.Db.C("M_Employees").FindId(task.PrimarySellerObjectID).One(&emp)
	if err1 == nil {
		taskGet.PrimarySeller = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = dal.mongo.Db.C("M_Employees").FindId(task.PrimaryOCObjectID).One(&emp)
	if err1 == nil {
		taskGet.PrimaryOC = emp.Name
	}
	otherExecutorsCount := len(task.OtherExecutorObjectIds)
	taskGet.OtherExecutors = make([]string, otherExecutorsCount, otherExecutorsCount)
	for index, value := range task.OtherExecutorObjectIds {
		emp = new(types.EmployeeName)
		err1 = dal.mongo.Db.C("M_Employees").FindId(value).One(&emp)
		if err1 == nil {
			taskGet.OtherExecutors[index] = *emp.Name
		}
	}

	product := new(types.ProductName)
	err1 = dal.mongo.Db.C("T_Products").FindId(task.ParentProductObjectID).One(&emp)
	if err1 == nil {
		taskGet.ParentProduct = product.Name
	}
	project := new(types.ProjectName)
	err1 = dal.mongo.Db.C("T_Projects").FindId(task.ParentProjectObjectID).One(&emp)
	if err1 == nil {
		taskGet.ParentProject = project.Name
	}
	return
}

// AddTask 定义
func (dal *TaskDAL) AddTask(taskPost types.Task_Post) (s map[string]string, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	task := new(types.Task)
	common.StructDeepCopy(taskPost, task)
	task.OID = bson.NewObjectId()
	dal.mongo.UseDB("local")

	err = dal.mongo.UseCollection("T_Tasks")
	if err != nil {
		return
	}

	dateString := time.Now().Format("20060102")
	var maxID *types.MaxID
	taskRegex := "^T" + dateString + "[0-9]{4}$"
	iter := dal.mongo.Collection.Find(bson.M{"id": bson.M{"$regex": taskRegex}}).Sort("-id").Iter()
	iter.Next(&maxID)
	var id string
	if maxID == nil {
		id = "T" + dateString + "0001"
	} else {
		maxNum, _ := strconv.Atoi((*maxID.ID)[9:])
		if maxNum >= 9999 {
			err = errors.New("今天新建的任务超出最大限度")
			return
		}
		id = fmt.Sprintf("T%s%04d", dateString, maxNum+1)
	}
	task.ID = &id

	err = dal.mongo.Collection.Insert(task)
	if err != nil && strings.Contains(err.Error(), "E11000 duplicate key error collection:") {
		return dal.AddTask(taskPost)
	} else if err != nil {
		return
	}

	s = make(map[string]string)
	s["Id"] = *task.ID

	return
}
