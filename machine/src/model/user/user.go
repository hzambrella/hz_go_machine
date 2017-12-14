package user

//用户基本信息
type UserBaseInfo struct {
	UserId     int64  `json:"user_id"`
	RoleCode   string `json:"role_code"`
	ParentId   int    `json:"parent_id"`
	ParentName string `json:"parent_name"`//用户名
	ParentRealName string `json:"parent_real_name"`//真实名
	//Password   string `json:"password"`
	Mobile     string `json:"mobile"`
	UserName   string `json:"user_name"`
	RealName   string `json:"real_name"`
	IdCard     string `json:"id_card"`
	BankCard   string `json:"bank_card"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

//用户金额信息
type UserBalanceInfo struct {
	UserId     int64
	//可提现佣金
	Cash       string
	//拥金余额(压款+可提现金额，单位:元后小数点4位)
	Balance    string
	//已提现佣金
	Withdrew   string
	//佣金总额
	Total      string
	CreateTime string
	UpdateTime string
}

//用户状态
const (
	Normal = 0
	//冻结
	Frozen = 1
	//审核中
	Checking = 2
)

//用户角色
const (
	ROOT     = "ROOT"
	ADMIN    = "ADMIN"
	CONSUMER = "CONSUMER"
)

//用户异常状态的提示语
const (
	FrozenMes   = "您的账户已被冻结"
	CheckingMes = "您的账户正在审核"
	NotFoundMes = "用户不存在"
	WrongPassMes="密码错误" 
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
