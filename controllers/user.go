package controllers

import (
	"TaskManagementSystem_Api/models"
	"TaskManagementSystem_Api/models/blls"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}
type signInUserStruct struct {
	UID      *string `json:"uid"`
	Password *string `json:"password"`
}
type ResponeBodyStruct struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// @Title Post_GetToken
// @Description get token by code
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 empty Token in Header
// @Failure 403 userinfo is empty
// @router user/token/ [post]
func (u *UserController) Post_GetToken() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	if token != "" {
		err := (&blls.UserBLL{}).ValidateToken(token)
		if err != nil {
			body.Error = err.Error()
		} else {
			body.Data = token
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
		token, err := (&blls.UserBLL{}).GetToken(*user.UID, *user.Password)
		if err != nil {
			body.Error = err.Error()
			// u.Ctx.Output.SetStatus(401)
		} else {
			body.Data = token
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

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
