package blls

import (
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
)

type TaskBLL struct {
}

func (bll TaskBLL) GetAllTasks() (u map[string]*types.Task, err error) {
	return (&dals.TaskDAL{}).GetHeaderOfAllTasks()
}
