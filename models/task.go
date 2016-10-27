package models

import (
	"TaskManagementSystem_Api/models/blls"
	"TaskManagementSystem_Api/models/types"
)

func GetAllTasks() (t []*types.TaskHeader, err error) {
	t, err = (&blls.TaskBLL{}).GetAllTasks()
	return
}

func GetTask(id string) (t types.Task, err error) {
	t, err = (&blls.TaskBLL{}).GetTaskDetail(id)
	return
}
