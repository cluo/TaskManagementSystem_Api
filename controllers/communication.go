package controllers

import (
	"TaskManagementSystem_Api/models"
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
	var communication types.Communication_Post
	json.Unmarshal(u.Ctx.Input.RequestBody, &communication)
	data, err := models.AddCommunication(communication)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = data
	}
	u.ServeJSON()
}

// @Title Get
// @Description get communication by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} types.Communication
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *CommunicationController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		communication, err := models.GetCommunications(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = communication
		}
	}
	u.ServeJSON()
}
