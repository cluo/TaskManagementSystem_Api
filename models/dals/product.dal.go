package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"
	"errors"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ProductDAL 定义
type ProductDAL struct {
	mongo *common.MongoSessionStruct
}

// GetProductHeaders 定义
func (dal *ProductDAL) GetProductHeaders(pageSize, pageNumber int) (productGetList []*types.ProductHeader_Get, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Products")
	if err != nil {
		return
	}

	if pageSize < 5 {
		pageSize = 5
	}
	if pageNumber < 1 {
		pageNumber = 1
	}
	var productList []*types.ProductHeader
	err = dal.mongo.Collection.Find(nil).Sort("-id").Skip((pageNumber - 1) * pageSize).Limit(pageSize).All(&productList)
	if err != nil {
		return
	}
	productCount := len(productList)
	productGetList = make([]*types.ProductHeader_Get, productCount, productCount)
	for index, value := range productList {
		productGet := new(types.ProductHeader_Get)
		common.StructDeepCopy(value, productGet)
		// 获取人员姓名
		emp := new(types.EmployeeName)
		if value.ProductManagerObjectID != nil {
			err1 := dal.mongo.Db.C("M_Employees").FindId(*value.ProductManagerObjectID).One(&emp)
			if err1 == nil {
				productGet.ProductManager = emp.Name
			}
		}
		productGetList[index] = productGet
	}
	return
}

// GetProductCount 定义
func (dal *ProductDAL) GetProductCount() (counts map[string]int, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Products")
	if err != nil {
		return
	}
	counts = make(map[string]int)
	// 所有产品数
	totalCount, err1 := dal.mongo.Collection.Find(nil).Count()
	if err1 != nil {
		err = err1
		return
	}
	counts["total"] = totalCount
	// 未开始产品数
	notStartedCount, err1 := dal.mongo.Collection.Find(bson.M{"status": "未开始"}).Count()
	if err1 != nil {
		err = err1
		return
	}
	counts["notStarted"] = notStartedCount
	// 进行中产品数
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	onGoingCount, err1 := dal.mongo.Collection.Find(bson.M{"status": "进行中", "planningEndDate": bson.M{"$lte": date}}).Count()
	if err1 != nil {
		err = err1
		return
	}
	counts["onGoing"] = onGoingCount
	// 超时产品数
	overtimeCount, err1 := dal.mongo.Collection.Find(bson.M{"status": "进行中", "planningEndDate": bson.M{"$gt": date}}).Count()
	if err1 != nil {
		err = err1
		return
	}
	counts["overtime"] = overtimeCount
	return
}

// GetProductDetail 定义
func (dal *ProductDAL) GetProductDetail(id string) (productGet *types.Product_Get, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Products")
	if err != nil {
		return
	}
	productGet = new(types.Product_Get)
	product := new(types.Product)
	dal.mongo.Collection.Find(bson.M{"id": id}).One(product)
	common.StructDeepCopy(product, productGet)

	// 获取人员姓名
	emp := new(types.EmployeeName)
	err1 := dal.mongo.Db.C("M_Employees").FindId(product.ProductManagerObjectID).One(&emp)
	if err1 == nil {
		productGet.ProductManager = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = dal.mongo.Db.C("M_Employees").FindId(product.CreatorObjectID).One(&emp)
	if err1 == nil {
		productGet.Creator = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = dal.mongo.Db.C("M_Employees").FindId(product.MarketingManagerObjectID).One(&emp)
	if err1 == nil {
		productGet.MarketingManager = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = dal.mongo.Db.C("M_Employees").FindId(product.DevelopmentManagerObjectID).One(&emp)
	if err1 == nil {
		productGet.DevelopmentManager = emp.Name
	}

	return
}

// AddProduct 定义
func (dal *ProductDAL) AddProduct(productPost types.Product_Post, user types.UserInfo_Get) (err error) {
	if !user.CheckPermissions(1, 11, 19) {
		err = errors.New("当前登录不能添加产品。")
		return
	}
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")

	err = dal.mongo.UseCollection("T_Products")
	if err != nil {
		return
	}

	product := new(types.Product)
	common.StructDeepCopy(productPost, product)
	product.RealReleaseDate = nil

	product.OID = bson.NewObjectId()
	now := time.Now()()
	product.CreatedTime = &now
	status := "新建"
	product.Status = &status
	objectID := new(types.ObjectID)
	product.CreatorID = user.EmpID
	err1 := dal.mongo.Db.C("M_Employees").Find(bson.M{"empId": product.CreatorID}).One(&objectID)
	if err1 != nil || objectID.Oid == nil {
		product.CreatorID = nil
	} else {
		product.CreatorObjectID = objectID.Oid
	}
	objectID = new(types.ObjectID)
	dal.mongo.Db.C("M_Employees").Find(bson.M{"empId": product.ProductManagerID}).One(&objectID)
	if err1 != nil || objectID.Oid == nil {
		product.ProductManagerID = nil
	} else {
		product.ProductManagerObjectID = objectID.Oid
	}
	objectID = new(types.ObjectID)
	dal.mongo.Db.C("M_Employees").Find(bson.M{"empId": product.MarketingManagerID}).One(&objectID)
	if err1 != nil || objectID.Oid == nil {
		product.MarketingManagerID = nil
	} else {
		product.MarketingManagerObjectID = objectID.Oid
	}
	objectID = new(types.ObjectID)
	dal.mongo.Db.C("M_Employees").Find(bson.M{"empId": product.DevelopmentManagerID}).One(&objectID)
	if err1 != nil || objectID.Oid == nil {
		product.DevelopmentManagerID = nil
	} else {
		product.DevelopmentManagerObjectID = objectID.Oid
	}
	err = dal.mongo.Collection.Insert(product)
	if err != nil && strings.Contains(err.Error(), "E11000 duplicate key error collection:") {
		return dal.AddProduct(productPost, user)
	} else if err != nil {
		return
	}

	return
}

// DeleteProduct 定义
func (dal *ProductDAL) DeleteProduct(id string, user types.UserInfo_Get) (err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Products")
	if err != nil {
		return
	}

	product := new(types.Product)
	err = dal.mongo.Collection.Find(bson.M{"id": id}).One(product)
	if err != nil {
		return
	}
	if (*product.Status != "新建" && *product.Status != "未开始" && *product.Status != "分配中" && *product.Status != "计划中") && !user.CheckPermissions(1) {
		err = errors.New("当前产品已经开始，不能删除当前记录。")
		return
	}
	if *product.CreatorID != *user.EmpID && !user.CheckPermissions(1) {
		err = errors.New("与产品创建者不是同一用户，不能删除当前记录。")
		return
	}
	err = dal.mongo.Collection.RemoveId(product.OID)
	if err != nil {
		return
	}
	dal.mongo.Db.C("T_Communications").RemoveAll(bson.M{"relevantId": id})
	return
}

// UpdateProduct 定义
func (dal *ProductDAL) UpdateProduct(id string, product types.Product_Post, user types.UserInfo_Get) (err error) {

	return
}
func (dal *ProductDAL) setUpdateBsonMap(product types.Product_Post) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})

	if len(m) == 0 {
		err = errors.New("没有任何修改内容！")
		return
	}
	return m, err
}
