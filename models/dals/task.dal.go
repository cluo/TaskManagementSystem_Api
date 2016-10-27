package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"
	"log"
)

type TaskDAL struct {
	mongo *common.MongoSessionStruct
}

func (t *TaskDAL) GetAllTaskHeaders() (u map[string]*types.TaskHeader, err error) {
	t.mongo, err = common.GetMongoSession()
	if err != nil {
		return nil, err
	}
	defer t.mongo.CloseSession()
	log.Println("--------------------------------------")
	log.Println("GetHeaderOfAllTasks")
	log.Println("--------------------------------------")
	return nil, nil
}
