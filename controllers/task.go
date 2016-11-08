package controllers

import (
	"TaskManagementSystem_Api/models"
	"TaskManagementSystem_Api/models/types"
	"encoding/json"

	"TaskManagementSystem_Api/models/common"

	"github.com/astaxie/beego"
)

// Operations about Tasks
type TaskController struct {
	beego.Controller
}

// @Title CreateTask
// @Description create tasks
// @Param	body		body 	types.Task_Post	true		"body for task content"
// @Success 200 {int} models.Task.Id
// @Failure 403 body is empty
// @router / [post]
func (u *TaskController) Post() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	_, err := (&common.AuthorizeStruct{}).ValidateAuthorize(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	var task types.Task_Post
	json.Unmarshal(u.Ctx.Input.RequestBody, &task)
	data, err := models.AddTask(task)
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = data
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Tasks (Header)
// @Success 200 {object} types.TaskHeader_Get
// @router / [get]
func (u *TaskController) GetAll() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	_, err := (&common.AuthorizeStruct{}).ValidateAuthorize(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	tasks, err := models.GetAllTasks()
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = tasks
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title Get
// @Description get task by tid
// @Param	tid		path 	string	true		"The key for staticblock"
// @Success 200 {object} types.Task_Get
// @Failure 403 :tid is empty
// @router /:tid [get]
func (u *TaskController) Get() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	_, err := (&common.AuthorizeStruct{}).ValidateAuthorize(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	tid := u.GetString(":tid")
	if tid != "" {
		task, err := models.GetTask(tid)
		if err != nil {
			body.Error = err.Error()
		} else {
			body.Data = task
		}
		u.Data["json"] = body
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the task
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	types.Task	true		"body for task content"
// @Success 200 {object} models.Task
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *TaskController) Put() {
	// uid := u.GetString(":uid")
	// if uid != "" {
	// 	var task models.Task
	// 	json.Unmarshal(u.Ctx.Input.RequestBody, &task)
	// 	uu, err := models.UpdateTask(uid, &task)
	// 	if err != nil {
	// 		u.Data["json"] = err.Error()
	// 	} else {
	// 		u.Data["json"] = uu
	// 	}
	// }
	// u.ServeJSON()
}

// @Title Delete
// @Description delete the task
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *TaskController) Delete() {
	// uid := u.GetString(":uid")
	// models.DeleteTask(uid)
	// u.Data["json"] = "delete success!"
	// u.ServeJSON()
}
