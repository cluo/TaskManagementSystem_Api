package models

import (
	"TaskManagementSystem_Api/models/blls"
	"TaskManagementSystem_Api/models/types"
)

func GetAllTasks() (u map[string]*types.TaskHeader, err error) {
	return (&blls.TaskBLL{}).GetAllTasks()
}
