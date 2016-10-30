package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"

	"errors"

	"gopkg.in/mgo.v2/bson"
)

type CommunicationDAL struct {
	mongo *common.MongoSessionStruct
}

func (c *CommunicationDAL) GetCommunications(id string) (communicationsGet map[string][]*types.Communication_Get, err error) {
	c.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer c.mongo.CloseSession()
	c.mongo.UseDB("local")
	err = c.mongo.UseCollection("T_Communications")
	if err != nil {
		return
	}

	communicationsGet = make(map[string][]*types.Communication_Get)

	var communications []*types.Communication
	c.mongo.Collection.Find(bson.M{"relevantId": id}).Sort("sentTime").All(&communications)
	communicationCount := len(communications)
	communicationsGet["data"] = make([]*types.Communication_Get, communicationCount, communicationCount)
	for index, value := range communications {
		communicationGet := new(types.Communication_Get)
		common.StructDeepCopy(value, communicationGet)
		// 获取人员姓名
		emp := new(types.EmployeeName)
		err1 := c.mongo.Db.C("M_Employees").FindId(value.PersonObjectID).One(&emp)
		if err1 == nil {
			communicationGet.PersonName = emp.Name
		}
		communicationsGet["data"][index] = communicationGet
	}
	return
}

func (c *CommunicationDAL) AddCommunication(communicationPost types.Communication_Post) (s map[string]map[string]string, err error) {
	c.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer c.mongo.CloseSession()
	communication := new(types.Communication)
	common.StructDeepCopy(communicationPost, communication)
	communication.OID = bson.NewObjectId()
	c.mongo.UseDB("local")
	objectID := new(types.ObjectID)
	err = c.mongo.Db.C("T_Tasks").Find(bson.M{"id": communication.RelevantID}).One(&objectID)
	err = c.mongo.Db.C("T_Project").Find(bson.M{"id": communication.RelevantID}).One(&objectID)
	err = c.mongo.Db.C("T_Product").Find(bson.M{"id": communication.RelevantID}).One(&objectID)
	if bson.ObjectId.Valid(*objectID.Oid) {
		communication.RelevantObjectID = objectID.Oid
	} else {
		err = errors.New("RelevantID无效。")
	}

	objectID = new(types.ObjectID)
	err = c.mongo.Db.C("M_Employees").Find(bson.M{"empId": communication.PersonID}).One(&objectID)
	communication.PersonObjectID = objectID.Oid

	err = c.mongo.UseCollection("T_Communications")
	if err != nil {
		return
	}
	err = c.mongo.Collection.Insert(communication)
	if err != nil {
		return
	}
	s = make(map[string]map[string]string)
	s["data"] = make(map[string]string)
	s["data"]["relevantId"] = *communication.RelevantID

	return
}
