package types

import "gopkg.in/mgo.v2/bson"

type UserInfo_Get struct {
	EmpID       *string `json:"empId"`
	Dept        *string `json:"dept"`
	Pre         *string `json:"pre"`
	Name        *string `json:"name"`
	Permissions int     `json:"permissions"`
}
type EmployeeInfo struct {
	OID          bson.ObjectId  `bson:"_id"`
	EmpID        *string        `bson:"empId"`
	DeptObjectID *bson.ObjectId `bson:"deptObjectId"`
	Pre          *string        `bson:"pre-"`
	Name         *string        `bson:"name"`
	Permissions  int            `bson:"permissions"`
}

type EmployeeOid struct {
	OID *bson.ObjectId `bson:"empObjectId"`
}

type DeptName struct {
	Name *string `bson:"name"`
}
