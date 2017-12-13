package routes

//用户相关
import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/contrib/sessions"
	"model/user"
	"tool/gintool"
	//"tool/inicfg"
	"interceptor/auth"
)

const (
	userDetailViewPath = "/machine/user/detail/view"
	userDetailDataPath = "/machine/user/detail/data"
)

func init() {
	r := gintool.Default()
	r.GET(userDetailViewPath, userDetailView)
	r.POST(userDetailDataPath, userDetailData)
}

// 用户详情页面
func userDetailView(c *gin.Context) {
	c.String(200, "ok")
}

// 用户详情数据
func userDetailData(c *gin.Context) {
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

	userDetail, err := db.GetUserById(uss.Uid)
	if err != nil {
		if err == user.UserDataNotFound {
			h["message"] = user.NotFoundMes
			c.JSON(500, h)
			return
		}

		h["message"] = dbWrong
		c.JSON(500, h)
	}

	h["userid"]=userDetail.UserId
	h["bankCard"]=userDetail.BankCard
	h["createTime"]=userDetail.CreateTime
	h["idCard"]=userDetail.IdCard
	h["mobile"]=userDetail.Mobile
	h["parentId"]=userDetail.ParentId
	h["realName"]=userDetail.RealName
	h["roleCode"]=userDetail.RoleCode
	h["Status"]=userDetail.Status
	h["updateTime"]=userDetail.UpdateTime
	c.JSON(200,h)
}
