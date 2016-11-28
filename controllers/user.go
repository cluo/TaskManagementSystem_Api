package controllers

import (
	"TaskManagementSystem_Api/models/blls"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}
type signInUserStruct struct {
	UID         *string `json:"uid"`
	Password    *string `json:"password"`
	NewPassword *string `json:"newpassword"`
}
type ResponeBodyStruct struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// @Title SignIn
// @Description get token by code
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 empty Token in Header
// @Failure 403 userinfo is empty
// @router user/token/ [post]
func (u *UserController) SignIn() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	if token != "" {
		userinfo, err := (&blls.UserBLL{}).ValidateToken(token)
		if err != nil {
			body.Error = err.Error()
		} else {
			body.Data = userinfo
		}
		u.Data["json"] = body
		u.ServeJSON()
		return
	}
	user := new(signInUserStruct)
	err := json.Unmarshal(u.Ctx.Input.RequestBody, user)
	if err != nil {
		body.Error = err.Error()
	} else if user.UID != nil && user.Password != nil {
		userinfo, err := (&blls.UserBLL{}).SignIn(*user.UID, *user.Password)
		if err != nil {
			body.Error = err.Error()
			// u.Ctx.Output.SetStatus(401)
		} else {
			body.Data = userinfo
		}
	}
	u.Data["json"] = body
	u.ServeJSON()
}

func (u *UserController) GetMyUserInfo() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	userinfo, err := (&blls.UserBLL{}).GetUserInfo(token)
	if err != nil {
		body.Error = err.Error()
		u.Ctx.Output.SetStatus(401)
	} else {
		body.Data = userinfo
	}
	u.Data["json"] = body
	u.ServeJSON()
	return
}

func (u *UserController) ChangePassword() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	userInToken, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}
	uid := u.GetString(":uid")
	if uid != "" {
		user := new(signInUserStruct)
		err := json.Unmarshal(u.Ctx.Input.RequestBody, user)
		if *userInToken.UID == uid {
			err = (&blls.UserBLL{}).ChangePassword(uid, *user.Password, *user.NewPassword)
			if err == nil {
				body.Data = "update success!"
			} else {
				body.Error = err.Error()
			}
		} else {
			body.Error = "用户登录错误，请重新登录。"
		}
		u.Data["json"] = body
	}
	u.ServeJSON()
}
