package controllers

import (
	"TaskManagementSystem_Api/models"

	"github.com/astaxie/beego"
)

// Operations about Token
type TokenController struct {
	beego.Controller
}

// @Title Get
// @Description get token by code
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 empty Token in Header
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *TokenController) Get() {
	code := u.GetString(":code")
	if token != "" {
		token, err := models.GetToken(code)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = nil
		}
	}
	u.ServeJSON()
}
