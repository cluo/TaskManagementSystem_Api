package blls

import (
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
)

type ProjectBLL struct {
}

func (bll *ProjectBLL) GetProjects(pageSize, pageNumber int) (t []*types.ProjectHeader_Get, err error) {
	t, err = (&dals.ProjectDAL{}).GetProjectHeaders(pageSize, pageNumber)
	return
}
func (bll *ProjectBLL) GetProjectCount() (counts map[string]int, err error) {
	counts, err = (&dals.ProjectDAL{}).GetProjectCount()
	return
}
func (bll *ProjectBLL) GetProjectDetail(id string) (t *types.Project_Get, err error) {
	t, err = (&dals.ProjectDAL{}).GetProjectDetail(id)
	return
}
func (bll *ProjectBLL) AddProject(projectPost types.Project_Post, user types.UserInfo_Get) (err error) {
	err = (&dals.ProjectDAL{}).AddProject(projectPost, user)
	return
}

func (bll *ProjectBLL) DeleteProject(id string, user types.UserInfo_Get) (err error) {
	err = (&dals.ProjectDAL{}).DeleteProject(id, user)
	return
}
func (bll *ProjectBLL) UpdateProject(id string, project types.Project_Post, user types.UserInfo_Get) (err error) {
	err = (&dals.ProjectDAL{}).UpdateProject(id, project, user)
	return
}
