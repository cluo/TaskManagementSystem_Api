package controllers

import (
	"TaskManagementSystem_Api/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Tasks
type TaskController struct {
	beego.Controller
}

// @Title CreateTask
// @Description create tasks
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {int} models.Task.Id
// @Failure 403 body is empty
// @router / [post]
func (u *TaskController) Post() {
	var task models.Task
	json.Unmarshal(u.Ctx.Input.RequestBody, &task)
	uid := models.AddTask(task)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Tasks
// @Success 200 {object} models.Task
// @router / [get]
func (u *TaskController) GetAll() {
	tasks := models.GetAllTasks()
	u.Data["json"] = tasks
	u.ServeJSON()
}

// @Title Get
// @Description get task by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Task
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *TaskController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		task, err := models.GetTask(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = task
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the task
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {object} models.Task
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *TaskController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var task models.Task
		json.Unmarshal(u.Ctx.Input.RequestBody, &task)
		uu, err := models.UpdateTask(uid, &task)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the task
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *TaskController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteTask(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
