package blls

import (
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
)

type TaskBLL struct {
}

func (bll TaskBLL) GetAllTasks() (t []*types.TaskHeader_Get, err error) {
	t, err = (&dals.TaskDAL{}).GetAllTaskHeaders()
	return
}

func (bll TaskBLL) GetTaskDetail(id string) (t *types.Task_Get, err error) {
	t, err = (&dals.TaskDAL{}).GetTaskDetail(id)
	return
}
func (bll TaskBLL) AddTask(taskPost types.Task_Post) (s map[string]string, err error) {
	s, err = (&dals.TaskDAL{}).AddTask(taskPost)
	return
}
