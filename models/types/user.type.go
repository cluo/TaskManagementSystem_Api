package types

type UserInfo_Get struct {
	Token       *string `json:"token"`
	EmpID       *string `json:"empId"`
	Dept        *string `json:"dept"`
	Pre         *string `json:"pre"`
	Name        *string `json:"name"`
	Permissions []int   `json:"permissions"`
}

type DeptName struct {
	Name *string `bson:"name"`
}
