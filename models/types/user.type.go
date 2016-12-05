package types

type UserInfo_Get struct {
	UID         *string `json:"uid"`
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

// 权限设定
// 1        系统管理员
// 11       产品经理管理员
// 17       市场经理
// 18       技术支持
// 19       产品经理
// 21       项目经理管理员
// 29       项目经理

// 38       技术经理
// 39       研发工程师

// 98       销售
// 99       OC
// 100      测试
