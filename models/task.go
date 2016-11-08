package models

import (
	"TaskManagementSystem_Api/models/blls"
	"TaskManagementSystem_Api/models/types"
)

func GetAllTasks() (t []*types.TaskHeader_Get, err error) {
	t, err = (&blls.TaskBLL{}).GetAllTasks()
	return
}

func GetTask(id string) (t *types.Task_Get, err error) {
	t, err = (&blls.TaskBLL{}).GetTaskDetail(id)
	return
}

func AddTask(taskPost types.Task_Post) (s map[string]string, err error) {
	s, err = (&blls.TaskBLL{}).AddTask(taskPost)
	return
}
