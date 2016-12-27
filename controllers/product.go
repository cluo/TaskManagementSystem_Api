package controllers

import (
	"TaskManagementSystem_Api/models/blls"
	"TaskManagementSystem_Api/models/types"
	"encoding/json"
	"errors"

	"strings"

	"github.com/astaxie/beego"
)

// Operations about Products
type ProductController struct {
	beego.Controller
}

// @Title CreateProduct
// @Description create products
// @Param	body		body 	types.Product_Post	true		"body for product content"
// @Success 200 {int} (&blls.ProductBLL{}).Product.Id
// @Failure 403 body is empty
// @router / [post]
func (u *ProductController) Post() {
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

	var product types.Product_Post
	json.Unmarshal(u.Ctx.Input.RequestBody, &product)
	err = (&blls.ProductBLL{}).AddProduct(product, user)
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = "insert success!"
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title GetList
// @Description get all Products (Header)
// @Success 200 {object} types.ProductHeader_Get
// @router / [get]
func (u *ProductController) GetList() {
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

	products, err := (&blls.ProductBLL{}).GetProducts(pageSize, pageNumber)
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = products
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title GetList
// @Description get all GetAll (Id/Name)
// @Success 200 {object} types.ProductHeader_Get
// @router / [get]
func (u *ProductController) GetAll() {
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

	products, err := (&blls.ProductBLL{}).GetAllProducts()
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = products
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title GetProductCount
// @Description get Product Count
// @Success 200 {object}
// @router / [get]
func (u *ProductController) GetProductCount() {
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

	counts, err := (&blls.ProductBLL{}).GetProductCount()
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = counts
	}
	u.Data["json"] = body
	u.ServeJSON()
}

// @Title Get
// @Description get product by tid
// @Param	tid		path 	string	true		"The key for staticblock"
// @Success 200 {object} types.Product_Get
// @Failure 403 :tid is empty
// @router /:tid [get]
func (u *ProductController) Get() {
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
		product, err := (&blls.ProductBLL{}).GetProductDetail(tid)
		if err != nil {
			body.Error = err.Error()
		} else {
			body.Data = product
		}
		u.Data["json"] = body
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the product
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	types.Product	true		"body for product content"
// @Success 200 {object} (&blls.ProductBLL{}).Product
// @Failure 403 :uid or :method is empty
// @router /:uid/:method [put]
func (u *ProductController) Put() {
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
		var product types.Product_Post
		json.Unmarshal(u.Ctx.Input.RequestBody, &product)
		switch method {
		case "update":
			err = (&blls.ProductBLL{}).UpdateProduct(tid, product, user)
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
// @Description delete the product
// @Param	tid		path 	string	true		"The tid you want to delete"
// @Success 200 delete success!
// @Failure 403 tid is empty
// @router /:tid [delete]
func (u *ProductController) Delete() {
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
	err = (&blls.ProductBLL{}).DeleteProduct(tid, user)
	if err == nil {
		body.Data = "delete success!"
	} else {
		body.Error = err.Error()
	}
	u.Data["json"] = body
	u.ServeJSON()
}
