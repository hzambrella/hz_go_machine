package user

import (
	"database/sql"
	"errors"
	"time"
	"tool/datastore"
	"tool/money"
)

//用户
var (
	// 封装错误，方便route层针对不同的错误进行不同的处理。如仅仅是查找结果为空和数据库系统错误的处理方式就不一样
	UserDataNotFound = errors.New("user data not found")
	UserWrongPass    = errors.New("user pass is not correct")
	UserFrozen       = errors.New("user is frozen")
	UserChecking     = errors.New("user is checking")
	formatTime       = "2006-01-02 15:04:05"
)

type UserDB interface {
	Close()
	// 检查用户状态
	CheckUserByName(userName string) (bool, error)
	// 登录校验,返回用户id
	LoginCheck(userName string, pass string) (bool, int64, error)
	// 根据用户id查询用户基本详情信息
	GetUserBaseInfoById(userid int64) (*UserBaseInfo, error)
	// 根据用户id查询用户佣金详情信息
	GetUserBalanceInfoById(userid int64) (*UserBalanceInfo, error)
	//查询直接上级信息
	GetParentNode(userid int64) (*UserBaseInfo, error)
	//查询下级用户,若role_code为空，就是下级的所有角色
	GetChildNodeInfo(userid int64, role_code string) ([]UserBaseInfo, error)

	// 用户注册
	CreateUser(parentid int64, user *UserBaseInfo, pass string) (int64, error)

	//修改用户基本信息,这个接口不会修改密码status,parentid,role_code
	UpdateUserBaseInfo(user *UserBaseInfo) error
	//修改用户佣金信息,这个接口不会修改status
	//UpdateUserBalanceInfo(user UserBaseInfo) error
	//修改用户status
	UpdateUserStatus(userid int64, status int) error
}

type userDB struct {
	*sql.DB
}

func NewUserDB() (UserDB, error) {
	db, err := datastore.LinkStore.GetDB("master")
	if err != nil {
		return nil, err
	}
	udb := &userDB{db}
	return udb, nil
}

func (db *userDB) Close() {
	db.Close()
}

// 检查用户状态
func (db *userDB) CheckUserByName(userName string) (bool, error) {
	var status int
	err := db.QueryRow(checkUserByNameSql, userName).Scan(
		&status,
	)
	//没数据，系统错误
	if err != nil {
		if err == sql.ErrNoRows {
			return false, UserDataNotFound
		} else {
			return false, err
		}
	}
	return checkUserStatus(status)
}

//登录校验 返回用户id
func (db *userDB) LoginCheck(userName string, pass string) (bool, int64, error) {
	var id int64
	var passFromDB string
	var status int
	err := db.QueryRow(checkLoginSql, userName).Scan(
		&id,
		&passFromDB,
		&status,
	)

	//没数据，系统错误
	if err != nil {
		if err == sql.ErrNoRows {
			return false, 0, UserDataNotFound
		} else {
			return false, 0, err
		}
	}
	// 密码错误
	if pass != passFromDB {
		return false, 0, UserWrongPass
	}
	var isNormal bool
	isNormal, err = checkUserStatus(status)
	return isNormal, id, err
}

//通过id获取用户基本详情信息
func (db *userDB) GetUserBaseInfoById(id int64) (*UserBaseInfo, error) {
	var createTime time.Time
	var updateTime time.Time
	user := UserBaseInfo{}
	err := db.QueryRow(getUserBaseInfoByIdSql, id).Scan(
		&user.UserId,
		&user.RoleCode,
		&user.ParentId,
		//&user.Password,
		&user.Mobile,
		&user.UserName,
		&user.RealName,
		&user.IdCard,
		&user.BankCard,
		&user.Status,
		&createTime,
		&updateTime,
		&user.ParentName,
		&user.ParentRealName,	
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, UserDataNotFound
		} else {
			return nil, err
		}
	}

	user.CreateTime = createTime.Format(formatTime)
	user.UpdateTime = updateTime.Format(formatTime)
	return &user, nil
}

// 获取用户佣金数据s
func (db *userDB) GetUserBalanceInfoById(userid int64) (*UserBalanceInfo, error) {
	var createTime time.Time
	var updateTime time.Time
	var balance int
	var cash int
	var withdrew int
	var total int
	user := UserBalanceInfo{}
	err := db.QueryRow(getUserBalanceInfoByIdSql, userid).Scan(
		&user.UserId,
		&balance,
		&cash,
		&withdrew,
		&total,
		//&user.Password,
		&createTime,
		&updateTime,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, UserDataNotFound
		} else {
			return nil, err
		}
	}

	user.CreateTime = createTime.Format(formatTime)
	user.UpdateTime = updateTime.Format(formatTime)
	user.Balance=money.New(float64(balance) / 10000).Format(2)
	user.Withdrew=money.New(float64(withdrew) / 10000).Format(2)
	user.Cash=money.New(float64(cash) / 10000).Format(2)
	user.Total=money.New(float64(total) / 10000).Format(2)
	return &user, nil

}

//用户注册
func (db *userDB) CreateUser(parentid int64, info *UserBaseInfo, pass string) (int64, error) {
	var createStatus int
	parent, err := db.GetUserBaseInfoById(parentid)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, UserDataNotFound
		} else {
			return 0, err
		}
	}

	if parent.RoleCode == ADMIN || parent.RoleCode == ROOT {
		createStatus = Normal
	} else {
		createStatus = Checking
	}

	result, err := db.Exec(createUserBaseInfoSql,
		parentid, pass, info.Mobile, info.UserName, info.RealName,
		info.RoleCode, info.IdCard, info.BankCard, createStatus)
	if err != nil {
		return 0, err
	}

	uid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	if createStatus == Normal {
		if err := db.createUserBalanceInfo(uid); err != nil {
			return 0, err
		}
	}
	return uid, nil
}

// 创键一个用户佣金信息表
func (db *userDB) createUserBalanceInfo(userid int64) error {
	_, err := db.Exec(createUserBalanceInfoSql, userid)
	if err != nil {
		return err
	}
	return nil

}

//查询直接上级的基本信息
func (db *userDB) GetParentNode(userid int64) (*UserBaseInfo, error) {
	var createTime time.Time
	var updateTime time.Time
	user := UserBaseInfo{}
	err := db.QueryRow(getParentNodeInfoSql, userid).Scan(
		&user.UserId,
		&user.RoleCode,
		&user.ParentId,
		//&user.Password,
		&user.Mobile,
		&user.UserName,
		&user.RealName,
		&user.IdCard,
		&user.BankCard,
		&user.Status,
		&createTime,
		&updateTime,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, UserDataNotFound
		} else {
			return nil, err
		}
	}
	user.CreateTime = createTime.Format(formatTime)
	user.UpdateTime = updateTime.Format(formatTime)
	return &user, nil
}

//查询下级,若role_code为空，就是所有角色
func (db *userDB) GetChildNodeInfo(userid int64, role_code string) ([]UserBaseInfo, error) {
	if role_code == "" {
		//若为空就是查询所有角色的下级
		row, err := db.Query(getAllChildNodeInfoSql, userid, userid)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, UserDataNotFound
			} else {
				return nil, err
			}
		}
		return getUserChildData(row)

	} else {
		//若不为空就是查询某个身份的下级
		row, err := db.Query(getChildNodeWithRoleSql, userid, userid, role_code)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, UserDataNotFound
			} else {
				return nil, err
			}
		}
		return getUserChildData(row)
	}
}

func getUserChildData(row *sql.Rows) ([]UserBaseInfo, error) {
	userList := []UserBaseInfo{}
	for row.Next() {
		user := UserBaseInfo{}
		var createTime time.Time
		var updateTime time.Time
		if err := row.Scan(
			&user.UserId,
			&user.RoleCode,
			&user.ParentId,
			//&user.Password,
			&user.Mobile,
			&user.UserName,
			&user.RealName,
			&user.IdCard,
			&user.BankCard,
			&user.Status,
			&createTime,
			&updateTime,
			&user.ParentName,
			&user.ParentRealName,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, UserDataNotFound
			} else {
				return nil, err
			}
		}
		user.CreateTime = createTime.Format(formatTime)
		user.UpdateTime = updateTime.Format(formatTime)
		userList = append(userList, user)
	}
	return userList, nil

}

//修改用户基本信息
func (db *userDB) UpdateUserBaseInfo(user *UserBaseInfo) error {
	_, err := db.Exec(updateUserBaseInfoSql,
		user.Mobile, user.UserName, user.RealName, user.IdCard, user.BankCard,
		user.UserId)
	if err != nil {
		return err
	}
	return nil
}

//修改用户佣金信息
/*
func (db *userDB) UpdateUserBalanceInfo(user UserBaseInfo) error {
	return nil
}
*/

//修改用户status
func (db *userDB) UpdateUserStatus(userid int64, status int) error {
	_, err := db.Exec(updateUserStatusSql,
		status, userid)
	if err != nil {
		return err
	}
	return nil
}
