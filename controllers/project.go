package controllers

import (
	"TaskManagementSystem_Api/models/blls"
	"TaskManagementSystem_Api/models/types"
	"encoding/json"
	"errors"

	"strings"

	"github.com/astaxie/beego"
)

// Operations about Projects
type ProjectController struct {
	beego.Controller
}

// @Title CreateProject
// @Description create projects
// @Param	body		body 	types.Project_Post	true		"body for project content"
// @Success 200 {int} (&blls.ProjectBLL{}).Project.Id
// @Failure 403 body is empty
// @router / [post]
func (u *ProjectController) Post() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	user, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	var project types.Project_Post
	json.Unmarshal(u.Ctx.Input.RequestBody, &project)
	err = (&blls.ProjectBLL{}).AddProject(project, user)
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = "insert success!"
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title GetList
// @Description get all Projects (Header)
// @Success 200 {object} types.ProjectHeader_Get
// @router / [get]
func (u *ProjectController) GetList() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	_, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}
	pageSize, _ := u.GetInt("pagesize", 5)
	pageNumber, _ := u.GetInt("page", 1)

	projects, err := (&blls.ProjectBLL{}).GetProjects(pageSize, pageNumber)
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = projects
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Projects (id/name)
// @Success 200 {object} types.ProjectHeader_Get
// @router / [get]
func (u *ProjectController) GetAll() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	_, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	projects, err := (&blls.ProjectBLL{}).GetAllProjects()
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = projects
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title GetProjectCount
// @Description get Project Count
// @Success 200 {object}
// @router / [get]
func (u *ProjectController) GetProjectCount() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	_, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	counts, err := (&blls.ProjectBLL{}).GetProjectCount()
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = counts
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title Get
// @Description get project by tid
// @Param	tid		path 	string	true		"The key for staticblock"
// @Success 200 {object} types.Project_Get
// @Failure 403 :tid is empty
// @router /:tid [get]
func (u *ProjectController) Get() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	_, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	tid := u.GetString(":tid")
	if tid != "" {
		project, err := (&blls.ProjectBLL{}).GetProjectDetail(tid)
		if err != nil {
			body.Error = err.Error()
		} else {
			body.Data = project
		}
		u.Data["json"] = body
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the project
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	types.Project	true		"body for project content"
// @Success 200 {object} (&blls.ProjectBLL{}).Project
// @Failure 403 :uid or :method is empty
// @router /:uid/:method [put]
func (u *ProjectController) Put() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	user, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	tid := u.GetString(":tid")
	method := strings.ToLower(u.GetString(":method"))
	if tid != "" && method != "" {
		var project types.Project_Post
		json.Unmarshal(u.Ctx.Input.RequestBody, &project)
		switch method {
		case "update":
			err = (&blls.ProjectBLL{}).UpdateProject(tid, project, user)
			break
		default:
			err = errors.New("method参数错误，该操作不存在。")
		}
		if err == nil {
			body.Data = "update success!"
		} else {
			body.Error = err.Error()
		}
		u.Data["json"] = body
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the project
// @Param	tid		path 	string	true		"The tid you want to delete"
// @Success 200 delete success!
// @Failure 403 tid is empty
// @router /:tid [delete]
func (u *ProjectController) Delete() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	user, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	tid := u.GetString(":tid")
	err = (&blls.ProjectBLL{}).DeleteProject(tid, user)
	if err == nil {
		body.Data = "delete success!"
	} else {
		body.Error = err.Error()
	}
	u.Data["json"] = body
	u.ServeJSON()
}
