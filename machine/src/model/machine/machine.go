package machine

//机器码列表
type Machine struct{
	//机器码
	MachineCode string
	//是否是当前拥有者
	IsOwner bool
	//当前拥有者
	OwnerId int
	//记录的创建时间
	CreateTime string
	//更新时间
	UpdateTime string
	//状态
	Status int
}

// 机器码转手记录，从自己开始
type MachineFlow struct{
	//机器码
	MachineCode string
	//发放者
	FromUserId int
	//分配者
	ToUserId int
	//记录的创建时间
	CreateTime string
	//更新时间
	UpdateTime string
	//状态
	Status int
}

//机器码订单交易记录
type SaleOrder struct{
	OrderId string
	SellerId int
	PurchaseId int
	//机器码
	MachineCode string
	//记录的创建时间
	CreateTime string
	//更新时间
	UpdateTime string
	//状态
	Status int
	//地址信息
	Addr string
	//电话信息
	Mobile string
	//备注
	Memo string
}

