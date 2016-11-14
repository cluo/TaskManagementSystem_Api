package controllers

import (
	"TaskManagementSystem_Api/models/blls"

	"github.com/astaxie/beego"
)

// Operations about Employees
type EmployeeController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Employees
// @Success 200 {object} models.Employee
// @router / [get]
func (e *EmployeeController) GetAll() {
	body := &ResponeBodyStruct{}
	token := e.Ctx.Input.Header("X-Auth-Token")
	_, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		e.Data["json"] = body
		e.Ctx.Output.SetStatus(401)
		e.ServeJSON()
		return
	}

	employees, err := (&blls.EmployeeBLL{}).GetAllEmployees()
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = employees
	}
	e.Data["json"] = body
	e.ServeJSON()
}
