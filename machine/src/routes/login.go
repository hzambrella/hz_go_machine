package routes

//用户登录相关
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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
	registerViewPath    = "/machine/register/view"
	registerPath        = "/machine/register/post"
)

func init() {
	r := gintool.Default()
	r.GET(loginViewPath, loginView)
	r.POST(CheckUserByNamePath, CheckUserByName)
	r.POST(CheckUserByPassPath, CheckUserByPass)
	r.GET(registerViewPath, registerView)
	r.POST(registerPath, register)
}

// 主页面
func loginView(c *gin.Context) {
	c.HTML(200, "machine/login",gin.H{})
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
		fmt.Println(paramWrongFormat("userName"))
		c.String(400, "请输入用户名")
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
			fmt.Println(err)
			c.String(500,sysWrong)
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
		fmt.Println(paramWrongFormat("userName"))
		c.String(400,"请输入用户名")
		return
	}

	passWord := FormValue(c, "passWord")
	if passWord == "" {
		fmt.Println(paramWrongFormat("passWord"))
		c.String(400, "请输入密码")
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
			fmt.Println(err)
			c.String(500,sysWrong)
			return
		}
		c.JSON(200, h)
		return
	}

	if isPass == true && err == nil {
		if err := auth.SetUserSession(uid, c); err != nil {
			fmt.Println(err)
			c.String(500,sysWrong)
			return
		}
		h["message"] = "用户有效"
		h["isValid"] = true
		c.JSON(200, h)
		return
	}
}

func registerView(c *gin.Context) {
	c.JSON(200, "ok")
}

func register(c *gin.Context) {
	db, err := user.NewUserDB()
	if err != nil {
		fmt.Println(dbWrong)
		c.String(500, dbWrong)
		return
	}
	// parentid
	parentidStr := FormValue(c, "parentid")
	if parentidStr == "" {
		fmt.Println(paramWrongFormat("parentid"))
		c.String(400, paramWrongFormat("parentid"))
		return
	}

	parentid, err := strconv.ParseInt(parentidStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		c.String(500, sysWrong)
		return
	}

	//密码
	passWord := FormValue(c, "passWord")
	if passWord == "" {
		fmt.Println(paramWrongFormat("passWord"))
		c.String(400, paramWrongFormat("密码"))
		return
	}
	//用户名
	userName := FormValue(c, "userName")
	if userName == "" {
		fmt.Println(paramWrongFormat("userName"))
		c.String(400, paramWrongFormat("用户名"))
		return
	}
	//真实名
	realName := FormValue(c, "realName")
	if realName == "" {
		fmt.Println(paramWrongFormat("realName"))
		c.String(400, paramWrongFormat("真实姓名"))
		return
	}
	//角色
	roleCode := FormValue(c, "roleCode")
	if roleCode == "" {
		fmt.Println(paramWrongFormat("roleCode"))
		c.String(400, paramWrongFormat("角色名"))
		return
	}
	//身份证号
	iDCard := FormValue(c, "idCard")
	if iDCard == "" {
		fmt.Println(paramWrongFormat("idCard"))
		c.String(400, paramWrongFormat("身份证号"))
		return
	}
	// 银行卡
	bankCard := FormValue(c, "bankCard")
	if bankCard == "" {
		fmt.Println(paramWrongFormat("bankCard"))
		c.String(400, paramWrongFormat("银行卡号"))
		return
	}
	// 手机号
	mobile := FormValue(c, "mobile")
	if bankCard == "" {
		fmt.Println(paramWrongFormat("mobile"))
		c.String(400, paramWrongFormat("手机号"))
		return
	}

	userR := &user.UserBaseInfo{
		Mobile:   mobile,
		UserName: userName,
		RealName: realName,
		RoleCode: roleCode,
		IdCard:   iDCard,
		BankCard: bankCard,
	}

	uid, err := db.CreateUser(parentid, userR, passWord)
	if err != nil {
		fmt.Println(err)
		c.String(500, sysWrong)
		return
	}

	c.JSON(200, gin.H{
		"userid": uid,
	})
}
