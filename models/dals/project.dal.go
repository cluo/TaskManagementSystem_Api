package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"
	"errors"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ProjectDAL 定义
type ProjectDAL struct {
	mongo *common.MongoSessionStruct
}

// GetAllProjects 定义
func (dal *ProjectDAL) GetAllProjects() (projectGetList []*types.ProjectName, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Projects")
	if err != nil {
		return
	}

	err = dal.mongo.Collection.Find(bson.M{"status": bson.M{"$gt": "已关闭"}}).Sort("-id").All(&projectGetList)
	return
}

// GetProjectHeaders 定义
func (dal *ProjectDAL) GetProjectHeaders(pageSize, pageNumber int) (projectGetList []*types.ProjectHeader_Get, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Projects")
	if err != nil {
		return
	}

	if pageSize < 5 {
		pageSize = 5
	}
	if pageNumber < 1 {
		pageNumber = 1
	}
	var projectList []*types.ProjectHeader
	err = dal.mongo.Collection.Find(nil).Sort("-id").Skip((pageNumber - 1) * pageSize).Limit(pageSize).All(&projectList)
	if err != nil {
		return
	}
	projectCount := len(projectList)
	projectGetList = make([]*types.ProjectHeader_Get, projectCount, projectCount)
	for index, value := range projectList {
		projectGet := new(types.ProjectHeader_Get)
		common.StructDeepCopy(value, projectGet)
		// 获取人员姓名
		emp := new(types.EmployeeName)
		if value.ProjectManagerObjectID != nil {
			err1 := dal.mongo.Db.C("M_Employees").FindId(*value.ProjectManagerObjectID).One(&emp)
			if err1 == nil {
				projectGet.ProjectManager = emp.Name
			}
		}
		projectGetList[index] = projectGet
	}
	return
}

// GetProjectCount 定义
func (dal *ProjectDAL) GetProjectCount() (counts map[string]int, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Projects")
	if err != nil {
		return
	}
	counts = make(map[string]int)
	// 所有项目数
	totalCount, err1 := dal.mongo.Collection.Find(nil).Count()
	if err1 != nil {
		err = err1
		return
	}
	counts["total"] = totalCount
	// 未开始项目数
	notStartedCount, err1 := dal.mongo.Collection.Find(bson.M{"status": "未开始"}).Count()
	if err1 != nil {
		err = err1
		return
	}
	counts["notStarted"] = notStartedCount
	// 进行中项目数
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	onGoingCount, err1 := dal.mongo.Collection.Find(bson.M{"status": "进行中", "planningEndDate": bson.M{"$lte": date}}).Count()
	if err1 != nil {
		err = err1
		return
	}
	counts["onGoing"] = onGoingCount
	// 超时项目数
	overtimeCount, err1 := dal.mongo.Collection.Find(bson.M{"status": "进行中", "planningEndDate": bson.M{"$gt": date}}).Count()
	if err1 != nil {
		err = err1
		return
	}
	counts["overtime"] = overtimeCount
	return
}

// GetProjectDetail 定义
func (dal *ProjectDAL) GetProjectDetail(id string) (projectGet *types.Project_Get, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Projects")
	if err != nil {
		return
	}
	projectGet = new(types.Project_Get)
	project := new(types.Project)
	dal.mongo.Collection.Find(bson.M{"id": id}).One(project)
	common.StructDeepCopy(project, projectGet)

	// 获取人员姓名
	emp := new(types.EmployeeName)
	err1 := dal.mongo.Db.C("M_Employees").FindId(project.CreatorObjectID).One(&emp)
	if err1 == nil {
		projectGet.Creator = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = dal.mongo.Db.C("M_Employees").FindId(project.PrimarySellerObjectID).One(&emp)
	if err1 == nil {
		projectGet.PrimarySeller = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = dal.mongo.Db.C("M_Employees").FindId(project.ProjectManagerObjectID).One(&emp)
	if err1 == nil {
		projectGet.ProjectManager = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = dal.mongo.Db.C("M_Employees").FindId(project.ProductManagerObjectID).One(&emp)
	if err1 == nil {
		projectGet.ProductManagerID = emp.Name
	}
	emp = new(types.EmployeeName)
	err1 = dal.mongo.Db.C("M_Employees").FindId(project.DevelopmentManagerObjectID).One(&emp)
	if err1 == nil {
		projectGet.DevelopmentManager = emp.Name
	}
	product := new(types.ProductName)
	err1 = dal.mongo.Db.C("T_Products").FindId(project.RelevantProductObjectID).One(&emp)
	if err1 == nil {
		projectGet.RelevantProduct = product.Name
	}
	return
}

// AddProject 定义
func (dal *ProjectDAL) AddProject(projectPost types.Project_Post, user types.UserInfo_Get) (err error) {
	if !user.CheckPermissions(1, 21, 29) {
		err = errors.New("当前登录不能添加项目。")
		return
	}
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")

	err = dal.mongo.UseCollection("T_Projects")
	if err != nil {
		return
	}

	project := new(types.Project)
	common.StructDeepCopy(projectPost, project)
	project.RealReleaseDate = nil

	project.OID = bson.NewObjectId()
	now := time.Now()
	project.CreatedTime = &now
	status := "新建"
	project.Status = &status
	objectID := new(types.ObjectID)
	project.CreatorID = user.EmpID
	err1 := dal.mongo.Db.C("M_Employees").Find(bson.M{"empId": project.CreatorID}).One(&objectID)
	if err1 != nil || objectID.Oid == nil {
		project.CreatorID = nil
	} else {
		project.CreatorObjectID = objectID.Oid
	}
	objectID = new(types.ObjectID)
	dal.mongo.Db.C("M_Employees").Find(bson.M{"empId": project.ProjectManagerID}).One(&objectID)
	if err1 != nil || objectID.Oid == nil {
		project.ProjectManagerID = nil
	} else {
		project.ProjectManagerObjectID = objectID.Oid
	}
	objectID = new(types.ObjectID)
	dal.mongo.Db.C("M_Employees").Find(bson.M{"empId": project.PrimarySellerID}).One(&objectID)
	if err1 != nil || objectID.Oid == nil {
		project.PrimarySellerID = nil
	} else {
		project.PrimarySellerObjectID = objectID.Oid
	}
	objectID = new(types.ObjectID)
	dal.mongo.Db.C("M_Employees").Find(bson.M{"empId": project.ProductManagerID}).One(&objectID)
	if err1 != nil || objectID.Oid == nil {
		project.ProductManagerID = nil
	} else {
		project.ProductManagerObjectID = objectID.Oid
	}
	objectID = new(types.ObjectID)
	dal.mongo.Db.C("M_Employees").Find(bson.M{"empId": project.DevelopmentManagerID}).One(&objectID)
	if err1 != nil || objectID.Oid == nil {
		project.DevelopmentManagerID = nil
	} else {
		project.DevelopmentManagerObjectID = objectID.Oid
	}
	err = dal.mongo.Collection.Insert(project)
	if err != nil && strings.Contains(err.Error(), "E11000 duplicate key error collection:") {
		return dal.AddProject(projectPost, user)
	} else if err != nil {
		return
	}

	return
}

// DeleteProject 定义
func (dal *ProjectDAL) DeleteProject(id string, user types.UserInfo_Get) (err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Projects")
	if err != nil {
		return
	}

	project := new(types.Project)
	err = dal.mongo.Collection.Find(bson.M{"id": id}).One(project)
	if err != nil {
		return
	}
	if (*project.Status != "新建" && *project.Status != "未开始" && *project.Status != "分配中" && *project.Status != "计划中") && !user.CheckPermissions(1) {
		err = errors.New("当前项目已经开始，不能删除当前记录。")
		return
	}
	if *project.CreatorID != *user.EmpID && !user.CheckPermissions(1) {
		err = errors.New("与项目创建者不是同一用户，不能删除当前记录。")
		return
	}
	err = dal.mongo.Collection.RemoveId(project.OID)
	if err != nil {
		return
	}
	dal.mongo.Db.C("T_Communications").RemoveAll(bson.M{"relevantId": id})
	return
}

// UpdateProject 定义
func (dal *ProjectDAL) UpdateProject(id string, project types.Project_Post, user types.UserInfo_Get) (err error) {

	return
}
func (dal *ProjectDAL) setUpdateBsonMap(project types.Project_Post) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})

	if len(m) == 0 {
		err = errors.New("没有任何修改内容！")
		return
	}
	return m, err
}
