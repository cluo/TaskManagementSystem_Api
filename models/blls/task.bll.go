package blls

import (
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
)

type TaskBLL struct {
	dal dals.TaskDAL
}

func (bll TaskBLL) GetAllTasks() (u map[string]*types.Task, err error) {
	return nil, nil
}
