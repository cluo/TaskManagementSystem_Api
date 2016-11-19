// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"TaskManagementSystem_Api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/v1/user", &controllers.UserController{}, "get:GetMyUserInfo")
	beego.Router("/v1/user/signin", &controllers.UserController{}, "post:SignIn")
	beego.Router("/v1/employee", &controllers.EmployeeController{}, "get:GetAll")

	beego.Router("/v1/task", &controllers.TaskController{}, "get:GetList;post:Post")
	beego.Router("/v1/task/counts", &controllers.TaskController{}, "get:GetTaskCount")
	beego.Router("/v1/task/:tid", &controllers.TaskController{}, "get:Get;delete:Delete;put:Put")

	beego.Router("/v1/communication/:id", &controllers.CommunicationController{}, "get:Get")
	beego.Router("/v1/communication", &controllers.CommunicationController{}, "post:Post")
}
