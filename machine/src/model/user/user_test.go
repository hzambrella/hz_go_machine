package user

import (
	//"database/sql"
	"fmt"
	"testing"
)

var db UserDB

func init() {
	var err error

	db, err = NewUserDB()
	if err != nil {
		panic(err)
	}
}

func TestUserCheck(t *testing.T) {
	isRight, err := db.CheckUserByName("root")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(isRight)
}

func TestUserLogin(t *testing.T) {
	isRight, id, err := db.LoginCheck("root", "root")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(id, isRight)

}

func TestUserDetail(t *testing.T) {
	user, err := db.GetUserById(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)

}
