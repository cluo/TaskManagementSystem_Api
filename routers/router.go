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
	beego.Router("/v1/user/:uid", &controllers.UserController{}, "put:ChangePassword")
	beego.Router("/v1/employee", &controllers.EmployeeController{}, "get:GetAll")

	beego.Router("/v1/task", &controllers.TaskController{}, "get:GetList;post:Post")
	beego.Router("/v1/task/counts", &controllers.TaskController{}, "get:GetTaskCount")
	beego.Router("/v1/task/:tid", &controllers.TaskController{}, "get:Get;delete:Delete")
	beego.Router("/v1/task/:tid/:method", &controllers.TaskController{}, "put:Put")

	beego.Router("/v1/product", &controllers.ProductController{}, "get:GetList;post:Post")
	beego.Router("/v1/product/counts", &controllers.ProductController{}, "get:GetProductCount")
	beego.Router("/v1/product/:tid", &controllers.ProductController{}, "get:Get;delete:Delete")
	beego.Router("/v1/product/:tid/:method", &controllers.ProductController{}, "put:Put")

	beego.Router("/v1/project", &controllers.ProjectController{}, "get:GetList;post:Post")
	beego.Router("/v1/project/counts", &controllers.ProjectController{}, "get:GetProjectCount")
	beego.Router("/v1/project/:tid", &controllers.ProjectController{}, "get:Get;delete:Delete")
	beego.Router("/v1/project/:tid/:method", &controllers.ProjectController{}, "put:Put")

	beego.Router("/v1/communication/:id", &controllers.CommunicationController{}, "get:Get")
	beego.Router("/v1/communication", &controllers.CommunicationController{}, "post:Post")
}
