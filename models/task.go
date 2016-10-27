package models

import (
	"TaskManagementSystem_Api/models/blls"
	"TaskManagementSystem_Api/models/types"
)

func GetAllTasks() (u map[string]*types.Task, err error) {
	return (&blls.TaskBLL{}).GetAllTasks()
}
