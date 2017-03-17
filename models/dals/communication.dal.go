package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"

	"errors"

	"gopkg.in/mgo.v2/bson"
)

// CommunicationDAL 定义
type CommunicationDAL struct {
	mongo *common.MongoSessionStruct
}

// GetCommunications 定义
func (dal *CommunicationDAL) GetCommunications(id string) (communicationsGet []*types.Communication_Get, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Communications")
	if err != nil {
		return
	}

	var communications []*types.Communication
	dal.mongo.Collection.Find(bson.M{"relevantId": id}).Sort("sentTime").All(&communications)
	communicationsCount := len(communications)
	communicationsGet = make([]*types.Communication_Get, communicationsCount, communicationsCount)
	for index, value := range communications {
		communicationGet := new(types.Communication_Get)
		common.StructDeepCopy(value, communicationGet)
		// 获取人员姓名
		emp := new(types.EmployeeName)
		err1 := dal.mongo.Db.C("M_Employees").FindId(value.PersonObjectID).One(&emp)
		if err1 == nil {
			communicationGet.PersonName = emp.Name
		}
		communicationsGet[index] = communicationGet
	}
	return
}

// AddCommunication 定义
func (dal *CommunicationDAL) AddCommunication(communicationPost types.Communication_Post) (s map[string]string, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	communication := new(types.Communication)
	common.StructDeepCopy(communicationPost, communication)
	communication.OID = bson.NewObjectId()
	dal.mongo.UseDB("local")
	objectID := new(types.ObjectID)
	err = dal.mongo.Db.C("T_Tasks").Find(bson.M{"id": communication.RelevantID}).One(&objectID)
	err = dal.mongo.Db.C("T_Projects").Find(bson.M{"id": communication.RelevantID}).One(&objectID)
	err = dal.mongo.Db.C("T_Products").Find(bson.M{"id": communication.RelevantID}).One(&objectID)
	if bson.ObjectId.Valid(*objectID.Oid) {
		communication.RelevantObjectID = objectID.Oid
	} else {
		err = errors.New("RelevantID无效。")
	}

	objectID = new(types.ObjectID)
	err = dal.mongo.Db.C("M_Employees").Find(bson.M{"empId": communication.PersonID}).One(&objectID)
	communication.PersonObjectID = objectID.Oid

	err = dal.mongo.UseCollection("T_Communications")
	if err != nil {
		return
	}
	err = dal.mongo.Collection.Insert(communication)
	if err != nil {
		return
	}
	s = make(map[string]string)
	s["relevantId"] = *communication.RelevantID

	return
}
