package controllers

import (
	"TaskManagementSystem_Api/models/blls"
	"TaskManagementSystem_Api/models/types"
	"encoding/json"
	"errors"
	"log"

	"strings"

	"github.com/astaxie/beego"
)

// Operations about Tasks
type TaskController struct {
	beego.Controller
}

// @Title CreateTask
// @Description create tasks
// @Param	body		body 	types.Task_Post	true		"body for task content"
// @Success 200 {int} (&blls.TaskBLL{}).Task.Id
// @Failure 403 body is empty
// @router / [post]
func (u *TaskController) Post() {
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

	var task types.Task_Post
	json.Unmarshal(u.Ctx.Input.RequestBody, &task)
	err = (&blls.TaskBLL{}).AddTask(task, user)
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = "insert success!"
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title GetList
// @Description get all Tasks (Header)
// @Success 200 {object} types.TaskHeader_Get
// @router / [get]
func (u *TaskController) GetList() {
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

	tasks, err := (&blls.TaskBLL{}).GetTasks(pageSize, pageNumber)
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = tasks
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title GetTaskCount
// @Description get Task Count
// @Success 200 {object}
// @router / [get]
func (u *TaskController) GetTaskCount() {
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

	counts, err := (&blls.TaskBLL{}).GetTaskCount()
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = counts
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
		task, err := (&blls.TaskBLL{}).GetTaskDetail(tid)
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
// @Success 200 {object} (&blls.TaskBLL{}).Task
// @Failure 403 :uid or :method is empty
// @router /:uid/:method [put]
func (u *TaskController) Put() {
	log.Println("controller update!")
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
		var task types.Task_Post
		json.Unmarshal(u.Ctx.Input.RequestBody, &task)
		switch method {
		case "update":
			err = (&blls.TaskBLL{}).UpdateTask(tid, task, user)
			break
		case "start":
			err = (&blls.TaskBLL{}).StartTask(tid, task, user)
			break
		case "finish":
			err = (&blls.TaskBLL{}).FinishTask(tid, task, user)
			break
		case "progress":
			err = (&blls.TaskBLL{}).ProgressTask(tid, task, user)
			break
		case "close":
			err = (&blls.TaskBLL{}).CloseTask(tid, task, user)
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
// @Description delete the task
// @Param	tid		path 	string	true		"The tid you want to delete"
// @Success 200 delete success!
// @Failure 403 tid is empty
// @router /:tid [delete]
func (u *TaskController) Delete() {
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
	err = (&blls.TaskBLL{}).DeleteTask(tid, user)
	if err == nil {
		body.Data = "delete success!"
	} else {
		body.Error = err.Error()
	}
	u.Data["json"] = body
	u.ServeJSON()
}
