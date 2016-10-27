package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"

	"gopkg.in/mgo.v2/bson"
)

type CommunicationDAL struct {
	mongo *common.MongoSessionStruct
}

func (t *CommunicationDAL) GetCommunications(id string) (communications map[string][]*types.Communication, err error) {
	t.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer t.mongo.CloseSession()
	t.mongo.UseDB("local")
	err = t.mongo.UseCollection("T_Communications")
	if err != nil {
		return
	}

	communications = make(map[string][]*types.Communication)
	communication := new(types.Communication)
	communications["data"] = make([]*types.Communication, 0, 10)

	iter := t.mongo.Collection.Find(bson.M{"relevantId": id}).Sort("sentTime").Iter()
	for iter.Next(&communication) {
		communications["data"] = append(communications["data"], communication)
		communication = new(types.Communication)
	}

	// 获取人员姓名
	for _, value := range communications["data"] {
		emp := new(types.EmployeeName)
		err1 := t.mongo.Db.C("M_Employees").FindId(value.PersonObjectID).One(&emp)
		if err1 == nil {
			value.PersonName = emp.Name
		}
	}
	return
}
