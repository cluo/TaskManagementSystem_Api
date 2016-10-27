package blls

import (
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
)

type TaskBLL struct {
}

func (bll TaskBLL) GetAllTasks() (t map[string][]*types.TaskHeader, err error) {
	t, err = (&dals.TaskDAL{}).GetAllTaskHeaders()
	return
}

func (bll TaskBLL) GetTaskDetail(id string) (t map[string]*types.Task, err error) {
	t, err = (&dals.TaskDAL{}).GetTaskDetail(id)
	return
}
