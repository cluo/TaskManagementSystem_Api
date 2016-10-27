package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"
)

type TaskDAL struct {
	mongo *common.MongoSessionStruct
}

func (t *TaskDAL) GetHeaderOfAllTasks() (u map[string]*types.Task, err error) {
	t.mongo, err = common.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer t.mongo.CloseSession()
	return nil, nil
}

// {
//     "_id" : ObjectId("580e02fdacfecf773c1ec350"),
//     "id" : "T201606120001",
//     "name" : "任务管理系统开发",
//     "resume" : "公安信息化产品中心任务管理系统开发",
//     "description" : "公安信息化产品中心任务管理系统开发，公安信息化产品中心任务管理系统开发，公安信息化产品中心任务管理系统开发。",
//     "customerContact" : "客户1 电话：13810138000",
//     "creatorId" : "000552",
//     "creatorObjectId" : ObjectId("580dec5bacfecf773c1ec327"),
//     "createdTime" : ISODate("2016-09-10T00:00:00.000Z"),
//     "primarySellerId" : "000186",
//     "primarySellerObjectId" : ObjectId("580df929acfecf773c1ec32f"),
//     "primaryOCId" : "000155",
//     "primaryOCObjectId" : ObjectId("580dee82acfecf773c1ec32d"),
//     "primaryExecutorId" : "000169",
//     "primaryExecutorObjectId" : ObjectId("580decd1acfecf773c1ec32a"),
//     "otherExecutorIds" : [
//         "000019",
//         "000800"
//     ],
//     "otherExecutorObjectIds" : [
//         ObjectId("580dec5bacfecf773c1ec328"),
//         ObjectId("580dfc4eacfecf773c1ec33a")
//     ],
//     "requiringBeginDate" : ISODate("2016-09-20T00:00:00.000Z"),
//     "requiringEndDate" : ISODate("2016-10-20T00:00:00.000Z"),
//     "planningBeginDate" : ISODate("2016-09-20T00:00:00.000Z"),
//     "planningEndDate" : ISODate("2016-10-20T00:00:00.000Z"),
//     "realBeginDate" : ISODate("2016-09-20T00:00:00.000Z"),
//     "realEndDate" : ISODate("2016-10-20T00:00:00.000Z"),
//     "percent" : 100,
//     "status" : "完成",
//     "parentProduct" : null,
//     "parentProject" : null
// }
