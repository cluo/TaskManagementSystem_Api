package types

type UserInfo_Get struct {
	Token       *string `json:"token"`
	EmpID       *string `json:"empId"`
	Dept        *string `json:"dept"`
	Pre         *string `json:"pre"`
	Name        *string `json:"name"`
	Permissions []int   `json:"permissions"`
}

func (u *UserInfo_Get) CheckPermissions(permissions ...int) bool {
	for _, value := range permissions {
		for _, p := range u.Permissions {
			if value == p {
				return true
			}
		}
	}
	return false
}

type DeptName struct {
	Name *string `bson:"name"`
}
