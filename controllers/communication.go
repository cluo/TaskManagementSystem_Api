package controllers

import (
	"TaskManagementSystem_Api/models/blls"
	"TaskManagementSystem_Api/models/types"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Communications
type CommunicationController struct {
	beego.Controller
}

// @Title AddCommunication
// @Description create communications
// @Param	body		body 	types.Communication_Post	true		"body for communication content"
// @Success 200 {string} types.communication.RelevantID
// @Failure 403 body is empty
// @router / [post]
func (u *CommunicationController) Post() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	var communication types.Communication_Post
	json.Unmarshal(u.Ctx.Input.RequestBody, &communication)
	data, err := (&blls.CommunicationBLL{}).AddCommunication(communication)
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = data
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title Get
// @Description get communication by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} types.Communication_Get
// @Failure 403 :id is empty
// @router /:id [get]
func (u *CommunicationController) Get() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	id := u.GetString(":id")
	if id != "" {
		communication, err := (&blls.CommunicationBLL{}).GetCommunications(id)
		if err != nil {
			body.Error = err.Error()
		} else {
			body.Data = communication
		}
	}
	u.Data["json"] = body
	u.ServeJSON()
}
