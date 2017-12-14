package user

import (
	//"database/sql"
	"fmt"
	"testing"
	"time"
)

var db UserDB

func init() {
	var err error

	db, err = NewUserDB()
	if err != nil {
		if err==UserDataNotFound{
			fmt.Println(UserDataNotFound.Error())
			return
		}
		panic(err)
	}
}

func TestCheckUserByName(t *testing.T) {
	isRight, err := db.CheckUserByName("root")
	if err != nil {
		if err==UserDataNotFound{
			fmt.Println(UserDataNotFound.Error())
			return
		}
		t.Fatal(err)
	}
	fmt.Println(isRight)
}

func TestLoginCheck(t *testing.T) {
	isRight, id, err := db.LoginCheck("root", "root")
	if err != nil {
		if err==UserDataNotFound{
			fmt.Println(UserDataNotFound.Error())
			return
		}
		t.Fatal(err)
	}
	fmt.Println(id, isRight)

}

func TestGetUserBaseInfoById(t *testing.T) {
	user, err := db.GetUserBaseInfoById(1)
	if err != nil {
		if err==UserDataNotFound{
			fmt.Println(UserDataNotFound.Error())
			return
		}
		t.Fatal(err)
	}
	fmt.Println(user)

}

func TestGetUserBalanceInfoById(t *testing.T) {
	user, err := db.GetUserBalanceInfoById(1)
	if err != nil {
		if err==UserDataNotFound{
			fmt.Println(UserDataNotFound.Error())
			return
		}
		t.Fatal(err)
	}
	fmt.Println(user)

}

func TestUserParentNode(t *testing.T) {
	user, err := db.GetParentNode(1)
	if err != nil {
		if err==UserDataNotFound{
			fmt.Println(UserDataNotFound.Error())
			return
		}
		t.Fatal(err)
	}
	fmt.Println(user)
}

func TestGetChildNodeInfo(t *testing.T) {
	list, err := db.GetChildNodeInfo(1, "CONSUMER")
	if err != nil {
		if err==UserDataNotFound{
			fmt.Println(UserDataNotFound.Error())
			return
		}
		t.Fatal(err)
	}
	fmt.Println(list)
}

/*
func TestCreateUser(t *testing.T){
	user:=&UserBaseInfo{
		//ParentId:100000,
		Mobile:"123132131",
		UserName:"hazhao",
		RealName:"hazho",
		RoleCode:CONSUMER,
		IdCard:"12",
		BankCard:"12",
	}
	uid,err:=db.CreateUser(1,user,"123456")
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println(uid)
}
*/

func TestUpdateUserBaseInfo(t *testing.T){
	user, err := db.GetUserBaseInfoById(1)
	if err != nil {
		if err==UserDataNotFound{
			fmt.Println(UserDataNotFound.Error())
			return
		}
		t.Fatal(err)
	}
	user.RealName="root"+time.Now().Format(formatTime)
	if err:=db.UpdateUserBaseInfo(user);err!=nil{
		t.Fatal(err)
	}
	fmt.Println("TestUpdateUserBaseInfo ok")
}

func TestUpdateUserStatus(t *testing.T){
	if err:=db.UpdateUserStatus(123213,Frozen);err!=nil{
		t.Fatal(err)
	}
	fmt.Println("TestUpdateUserStatus ok")
}