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

func (c *CommunicationDAL) GetCommunications(id string) (communications map[string][]*types.Communication, err error) {
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

	communications = make(map[string][]*types.Communication)
	communication := new(types.Communication)
	communications["data"] = make([]*types.Communication, 0, 10)

	iter := c.mongo.Collection.Find(bson.M{"relevantId": id}).Sort("sentTime").Iter()
	for iter.Next(&communication) {
		communications["data"] = append(communications["data"], communication)
		communication = new(types.Communication)
	}

	// 获取人员姓名
	for _, value := range communications["data"] {
		emp := new(types.EmployeeName)
		err1 := c.mongo.Db.C("M_Employees").FindId(value.PersonObjectID).One(&emp)
		if err1 == nil {
			value.PersonName = emp.Name
		}
	}
	return
}

func (c *CommunicationDAL) AddCommunication(communication types.Communication_Post) (s map[string]map[string]string, err error) {
	c.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer c.mongo.CloseSession()
	c.mongo.UseDB("local")
	objectID := new(types.ObjectID)
	err = c.mongo.Db.C("T_Tasks").Find(bson.M{"id": communication.RelevantID}).One(&objectID)
	err = c.mongo.Db.C("T_Project").Find(bson.M{"id": communication.RelevantID}).One(&objectID)
	err = c.mongo.Db.C("T_Product").Find(bson.M{"id": communication.RelevantID}).One(&objectID)
	if bson.ObjectId.Valid(objectID.Oid) {
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
	s["data"]["relevantId"] = communication.RelevantID

	return
}
