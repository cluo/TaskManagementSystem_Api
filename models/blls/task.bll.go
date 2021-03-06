package blls

import (
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
)

type TaskBLL struct {
}

func (bll *TaskBLL) GetTasks(pageSize, pageNumber int, searchCriteria, searchCriteria2 string, user types.UserInfo_Get) (t []*types.TaskHeader_Get, err error) {
	t, err = (&dals.TaskDAL{}).GetTaskHeaders(pageSize, pageNumber, searchCriteria, searchCriteria2, user)
	return
}
func (bll *TaskBLL) GetTaskCount() (counts map[string]int, err error) {
	counts, err = (&dals.TaskDAL{}).GetTaskCount()
	return
}
func (bll *TaskBLL) GetTaskDetail(id string) (t *types.Task_Get, err error) {
	t, err = (&dals.TaskDAL{}).GetTaskDetail(id)
	return
}
func (bll *TaskBLL) AddTask(taskPost types.Task_Post, user types.UserInfo_Get) (err error) {
	err = (&dals.TaskDAL{}).AddTask(taskPost, user)
	return
}

func (bll *TaskBLL) DeleteTask(id string, user types.UserInfo_Get) (err error) {
	err = (&dals.TaskDAL{}).DeleteTask(id, user)
	return
}
func (bll *TaskBLL) UpdateTask(id string, task types.Task_Post, user types.UserInfo_Get) (err error) {
	err = (&dals.TaskDAL{}).UpdateTask(id, task, user)
	return
}
func (bll *TaskBLL) StartTask(id string, task types.Task_Post, user types.UserInfo_Get) (err error) {
	err = (&dals.TaskDAL{}).StartTask(id, task, user)
	return
}
func (bll *TaskBLL) ProgressTask(id string, task types.Task_Post, user types.UserInfo_Get) (err error) {
	err = (&dals.TaskDAL{}).ProgressTask(id, task, user)
	return
}
func (bll *TaskBLL) FinishTask(id string, task types.Task_Post, user types.UserInfo_Get) (err error) {
	err = (&dals.TaskDAL{}).FinishTask(id, task, user)
	return
}
func (bll *TaskBLL) CloseTask(id string, task types.Task_Post, user types.UserInfo_Get) (err error) {
	err = (&dals.TaskDAL{}).CloseTask(id, task, user)
	return
}
func (bll *TaskBLL) RefuseTask(id string, task types.Task_Post, user types.UserInfo_Get) (err error) {
	err = (&dals.TaskDAL{}).RefuseTask(id, task, user)
	return
}
func (bll *TaskBLL) GetTaskScreen(pageSize, pageNumber int, typeString string) (t []*types.TaskScreen_Get, err error) {
	t, err = (&dals.TaskDAL{}).GetTaskScreen(pageSize, pageNumber, typeString)
	return
}
func (bll *TaskBLL) GetTaskScreenCount() (counts map[string]map[string]int, err error) {
	counts, err = (&dals.TaskDAL{}).GetTaskScreenCount()
	return
}
