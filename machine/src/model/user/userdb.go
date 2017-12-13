package user

import (
	"time"
	"database/sql"
	"errors"
	"tool/datastore"
)

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
	LoginCheck(userName string, pass string) (bool, int, error)
	GetUserById(id int) (*User, error)
	//GetUserByName(name string) (*User, error)
	//AddUser(name, password string) (int, error)
	//UpdateUserStatus(name, status string) (int, error)
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
func (db *userDB) LoginCheck(userName string, pass string) (bool, int, error) {
	var id int
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

//通过id获取用户详情
func (db *userDB) GetUserById(id int) (*User, error) {
	var createTime time.Time
	var updateTime time.Time
	user := User{}
	err := db.QueryRow(getUserByIdSql, id).Scan(
		&user.UserId,
		&user.RoleCode,
		&user.ParentId,
		&user.Password,
		&user.Mobile,
		&user.UserName,
		&user.RealName,
		&user.IdCard,
		&user.BankCard,
		&user.Status,
		&createTime,
		&updateTime,
	)
	user.CreateTime = createTime.Format(formatTime)
	user.UpdateTime = updateTime.Format(formatTime)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, UserDataNotFound
		} else {
			return nil, err
		}
	}
	return &user, nil

}
