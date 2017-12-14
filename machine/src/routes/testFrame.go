package routes

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/contrib/sessions"
	"interceptor/auth"
	"tool/gintool"
)

//框架测试
const (
	testJSONPath     = "/machine/json"
	testTemplatePath = "/machine/template"
	IndexViewPath    = "/machine/testdb"
)

func init() {
	r := gintool.Default()

	r.GET(testTemplatePath, testTemplate)
	r.GET(testJSONPath, testJSON)
	r.GET(IndexViewPath, IndexView)
}

// 主页面
func IndexView(c *gin.Context) {
	c.String(200, "ok")

}

func testJSON(c *gin.Context) {
	c.JSON(200, "test ok")

}

func testTemplate(c *gin.Context) {
	if !auth.Auth(c){
		return
	}
	
	userSess, err := auth.GetUserSession(c)
	if err != nil {
		c.String(500, err.Error())
	}
	c.HTML(200, "machine/test", gin.H{
		"map": userSess.Uid,
	})
}
