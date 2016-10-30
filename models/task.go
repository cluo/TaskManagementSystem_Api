package models

import (
	"TaskManagementSystem_Api/models/blls"
	"TaskManagementSystem_Api/models/types"
)

func GetAllTasks() (t map[string][]*types.TaskHeader_Get, err error) {
	t, err = (&blls.TaskBLL{}).GetAllTasks()
	return
}

func GetTask(id string) (t map[string]*types.Task_Get, err error) {
	t, err = (&blls.TaskBLL{}).GetTaskDetail(id)
	return
}
