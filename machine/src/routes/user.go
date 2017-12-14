package routes

//用户相关
import (
	"strconv"
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/contrib/sessions"
	"model/user"
	"tool/gintool"
	//"tool/inicfg"
	"interceptor/auth"
)

//用户相关
const (
	//用户详情
	userDetailViewPath = "/machine/user/detailView"
	userDetailDataPath = "/machine/user/detailData"
	//子级的列表
	getChildDataPath = "/machine/childData"
	//修改用户基本信息,这个接口不会修改密码,status,parentid,role_code
	updateUserStatusPath = "/machine/user/updateStatus"
	//修改用户status
	updateUserBaseInfoPath = "/machine/user/updateBaseInfo"
)

func init() {
	r := gintool.Default()
	r.GET(userDetailViewPath, userDetailView)
	r.GET(userDetailDataPath, userDetailData)
	r.GET(getChildDataPath, getChildData)
	r.POST(updateUserStatusPath, updateUserStatus)
	r.POST(updateUserBaseInfoPath, updateUserBaseInfo)
}

// 用户详情页面
func userDetailView(c *gin.Context) {
	if !auth.Auth(c){
		return
	}
	c.String(200, "ok")
}

// 用户详情数据
func userDetailData(c *gin.Context) {
	if !auth.Auth(c){
		return
	}

	h := gin.H{}
	uss, err := auth.GetUserSession(c)
	if err != nil {
		fmt.Println(err)
		h["message"] = sysWrong
		c.JSON(500, h)
		return
	}

	db, err := user.NewUserDB()
	if err != nil {
		fmt.Println(dbWrong)
		c.String(500, dbWrong)
		return
	}
	//用户基本详情
	userDetail, err := db.GetUserBaseInfoById(uss.Uid)
	if err != nil {
		if err == user.UserDataNotFound {
			h["message"] = user.NotFoundMes
			c.JSON(500, h)
			return
		}

		h["message"] = dbWrong
		c.JSON(500, h)
	}
	// 用户佣金详情
	userMoney, err := db.GetUserBalanceInfoById(uss.Uid)
	if err != nil {
		if err == user.UserDataNotFound {
			h["message"] = user.NotFoundMes
			c.JSON(500, h)
			return
		}

		h["message"] = dbWrong
		c.JSON(500, h)
	}

	//用户基本信息
	h["userid"] = userDetail.UserId
	h["bankCard"] = userDetail.BankCard
	h["createTime"] = userDetail.CreateTime
	h["idCard"] = userDetail.IdCard
	h["mobile"] = userDetail.Mobile
	h["realName"] = userDetail.RealName
	h["roleCode"] = userDetail.RoleCode
	h["Status"] = userDetail.Status
	h["updateTime"] = userDetail.UpdateTime
	//上级信息
	h["parentid"] = userDetail.ParentId
	h["parentName"] = userDetail.ParentName
	h["parentRealName"] = userDetail.ParentRealName
	//用户佣金信息
	h["balance"] = userMoney.Balance
	h["cash"] = userMoney.Cash
	h["total"] = userMoney.Total
	h["withdrew"] = userMoney.Withdrew
	c.JSON(200, h)
}

//获得子级的数据列表
func getChildData(c *gin.Context) {
	if !auth.Auth(c){
		return
	}

	h := gin.H{}
	uss, err := auth.GetUserSession(c)
	if err != nil {
		fmt.Println(err)
		h["message"] = sysWrong
		c.JSON(500, h)
		return
	}

	//角色
	roleCode := FormValue(c, "roleCode")

	db, err := user.NewUserDB()
	if err != nil {
		fmt.Println(dbWrong)
		c.String(500, dbWrong)
		return
	}

	uList, err := db.GetChildNodeInfo(uss.Uid, roleCode)
	if err != nil {
		if err == user.UserDataNotFound {
			h["message"] = user.NotFoundMes
			c.JSON(500, h)
			return
		}

		h["message"] = dbWrong
		c.JSON(500, h)
	}
	resultList := make([]map[string]interface{}, 0)

	for _, v := range uList {
		result := make(map[string]interface{}, 0)
		result["userid"] = v.UserId
		result["roleCode"] = v.RoleCode
		result["mobile"] = v.Mobile
		result["userName"] = v.UserName
		result["realName"] = v.RealName
		result["idCard"] = v.IdCard
		result["bankCard"] = v.BankCard
		result["status"] = v.Status
		result["createTime"] = v.CreateTime
		result["updateTime"] = v.UpdateTime
		result["createTime"] = v.CreateTime
		result["parentName"] = v.ParentName
		result["parentRealName"] = v.ParentRealName

		resultList = append(resultList, result)
	}
	h["resultList"] = resultList
	c.JSON(200, h)
}

//修改用户状态
func updateUserStatus(c *gin.Context){
	if !auth.Auth(c){
		return
	}
	uidStr:=FormValue(c,"userid")
	uid,err:=strconv.ParseInt(uidStr,10,64)
	if err!=nil{
		fmt.Println(err)
		c.String(500,paramWrongFormat("uid"))
	}

	statusStr:=FormValue(c,"status")
	status,err:=strconv.Atoi(statusStr)
	if err!=nil{
		fmt.Println(err)
		c.String(500,paramWrongFormat("status"))
	}
	db, err := user.NewUserDB()
	if err != nil {
		fmt.Println(dbWrong)
		c.String(500, dbWrong)
		return
	}

	if err:=db.UpdateUserStatus(uid,status);err!=nil{
		fmt.Println(err)
		c.String(500,dbWrong)
		return
	}
	c.String(200,success)
}

//修改用户信息
func updateUserBaseInfo(c *gin.Context){
	if !auth.Auth(c){
		return
	}

	uidStr:=FormValue(c,"userid")
	uid,err:=strconv.ParseInt(uidStr,10,64)
	if err!=nil{
		fmt.Println(err)
		c.String(500,paramWrongFormat("uid"))
	}

	db, err := user.NewUserDB()
	if err != nil {
		fmt.Println(dbWrong)
		c.String(500, dbWrong)
		return
	}
	u,err:=db.GetUserBaseInfoById(uid)
	if err!=nil{
		if err==user.UserDataNotFound{
			c.String(400,user.NotFoundMes)
			return
		}else{
			fmt.Println(err)
			c.String(500,sysWrong)
			return
		}
	}


	//真实名
	realName := FormValue(c, "realName")
	if realName != "" {
		u.RealName=realName
	}

	//身份证号
	iDCard := FormValue(c, "idCard")
	if iDCard != "" {
		u.IdCard=iDCard
	}

	// 银行卡
	bankCard := FormValue(c, "bankCard")
	if bankCard != "" {
		u.BankCard=bankCard
	}

	// 手机号
	mobile := FormValue(c, "mobile")
	if bankCard != "" {
		u.Mobile=mobile
	}

	if err:=db.UpdateUserBaseInfo(u);err!=nil{
		fmt.Println(err)
		c.String(500,dbWrong)
		return
	}
	c.String(200,success)
}