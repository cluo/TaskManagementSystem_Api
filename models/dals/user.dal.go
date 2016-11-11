package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// UserDAL 定义
type UserDAL struct {
	mongo *common.MongoSessionStruct
}

func (dal *UserDAL) GetUserInfo(uid string, password *string) (u *types.UserInfo_Get, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("M_Users")
	if err != nil {
		return
	}

	oid := new(types.EmployeeOid)
	if password != nil {
		dal.mongo.Collection.Find(bson.M{"username": uid}).One(oid)
	} else {
		dal.mongo.Collection.Find(bson.M{"username": uid, "password": *password}).One(oid)
	}
	if oid.OID == nil {
		err = errors.New("该用户在任务管理系统用户表中不存在，请联系管理员。")
		return
	}

	err = dal.mongo.UseCollection("M_Employees")
	if err != nil {
		return
	}

	empInfo := new(types.EmployeeInfo)
	dal.mongo.Collection.FindId(*oid.OID).One(empInfo)
	if err != nil {
		err = errors.New("该用户在任务管理系统雇员表中不存在，请联系管理员。")
		return
	}

	userInfo := new(types.UserInfo_Get)
	common.StructDeepCopy(empInfo, userInfo)
	dept := new(types.DeptName)
	err1 := dal.mongo.Db.C("M_Departments").FindId(empInfo.DeptObjectID).One(&dept)
	if err1 == nil {
		userInfo.Dept = dept.Name
	}
	u = userInfo
	return
}
