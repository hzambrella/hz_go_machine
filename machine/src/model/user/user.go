package user

type User struct {
	UserId     int    `json:"user_id"`
	RoleCode   string `json:"role_code"`
	ParentId   int    `json:"parent_id"`
	Password   string `json:"password"`
	Mobile     string `json:"mobile"`
	UserName   string `json:"user_name"`
	RealName   string `json:"real_name"`
	IdCard     string `json:"id_card"`
	BankCard   string `json:"bank_card"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

//用户状态
const (
	Normal = 0
	//冻结
	Frozen = 1
	//审核中
	Checking = 2
)

//用户异常状态的提示语
const (
	FrozenMes   = "您的账户已被冻结"
	CheckingMes = "您的账户正在审核"
	NotFoundMes="用户不存在"

)

//校验用户状态
func checkUserStatus(status int) (bool, error) {
	switch status {
	case Frozen:
		return false, UserFrozen
	case Checking:
		return false, UserChecking
	}
	return true, nil

}
