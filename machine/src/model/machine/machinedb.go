package machine

import (
//	"time"
	"database/sql"
	"errors"
	"tool/datastore"
)

//分销
var (
	// 封装错误，方便route层针对不同的错误进行不同的处理。如仅仅是查找结果为空和数据库系统错误的处理方式就不一样
	UserDataNotFound = errors.New("user data not found")
	UserWrongPass    = errors.New("user pass is not correct")
	UserFrozen       = errors.New("user is frozen")
	UserChecking     = errors.New("user is checking")
	formatTime       = "2006-01-02 15:04:05"
)

type MachineDB interface {
	Close()

	//机器码列表
	//所有机器码
	//已拥有机器码列表
	//已分配机器码列表
	//机器码动向查询
	GetMachineByUserID(userid int,operation string)
	//机器码分配
	DistributeMachine()
	//机器码下单
	SaleMachine()
	//机器码订单状态更改，当成功售出时触发分润
	UpdateMachineOrderStatus()
	
	//机器码销售记录
	GetMachineOrderList()
	//机器码销售记录详情
	GetMachineOrderDetail()
	
}

type machineDB struct {
	*sql.DB
}

func NewmachineDB() (machineDB, error) {
	db, err := datastore.LinkStore.GetDB("master")
	if err != nil {
		return nil, err
	}
	udb := &machineDB{db}
	return udb, nil
}

func (db *machineDB) Close() {
	db.Close()
}
