package routes

//用户登录相关
import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/contrib/sessions"
	"model/user"
	"tool/gintool"
	//"tool/inicfg"
	"interceptor/auth"
)

//将路由加入到 $ETCDIR/interceptor.ini中，不进行拦截和鉴权！！！！
const (
	loginViewPath       = "/machine/login/view"
	CheckUserByNamePath = "/machine/check/user"
	CheckUserByPassPath = "/machine/check/pass"
)

func init() {
	r := gintool.Default()
	r.GET(loginViewPath, loginView)
	r.POST(CheckUserByNamePath, CheckUserByName)
	r.POST(CheckUserByPassPath, CheckUserByPass)
}

// 主页面
func loginView(c *gin.Context) {
	c.String(200, "ok")
}

func CheckUserByName(c *gin.Context) {
	db, err := user.NewUserDB()
	if err != nil {
		fmt.Println(dbWrong)
		c.String(500, dbWrong)
		return
	}
	userName := FormValue(c, "userName")
	if userName == "" {
		fmt.Println(paramWrong)
		c.String(400, paramWrong)
		return
	}
	var isPass bool
	h := gin.H{}

	isPass, err = db.CheckUserByName(userName)
	if err != nil {
		h["isValid"] = false
		switch err {
		case user.UserDataNotFound:
			h["message"] = user.NotFoundMes
		case user.UserFrozen:
			h["message"] = user.FrozenMes
		case user.UserChecking:
			h["message"] = user.CheckingMes
		default:
			h["message"] = sysWrong
			c.JSON(500, h)
			return
		}
		c.JSON(200, h)
		return
	}

	if isPass == true && err == nil {
		h["message"] = "用户有效"
		h["isValid"] = true
		c.JSON(200, h)
	}
}

func CheckUserByPass(c *gin.Context) {
	db, err := user.NewUserDB()
	if err != nil {
		fmt.Println(dbWrong)
		c.String(500, dbWrong)
		return
	}
	userName := FormValue(c, "userName")
	if userName == "" {
		fmt.Println(paramWrong)
		c.String(400, paramWrong)
		return
	}

	passWord := FormValue(c, "passWord")
	if passWord == "" {
		fmt.Println(paramWrong)
		c.String(400, paramWrong)
		return
	}

	isPass, uid, err := db.LoginCheck(userName, passWord)
	h := gin.H{}

	if err != nil {
		h["isValid"] = false
		switch err {
		case user.UserDataNotFound:
			h["message"] = user.NotFoundMes
		case user.UserFrozen:
			h["message"] = user.FrozenMes
		case user.UserChecking:
			h["message"] = user.CheckingMes
		default:
			h["message"] = sysWrong
			c.JSON(500, h)
			return
		}
		c.JSON(200, h)
		return
	}

	if isPass == true && err == nil {
		if err := auth.SetUserSession(uid, c); err != nil {
			h["isValid"] = false
			h["message"] = sysWrong
			c.JSON(500, h)
		}
		h["message"] = "用户有效"
		h["isValid"] = true
		c.JSON(200, h)
		return
	}
}
