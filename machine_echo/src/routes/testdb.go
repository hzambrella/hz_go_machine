package routes

import (
	"github.com/labstack/echo"
	"tool/echotool"
)

const (
	testJSONPath     = "/machine/json"
	testTemplatePath = "/machine/template"
	IndexViewPath    = "/machine/testdb"
)


func init() {
	r := echotool.Default()

	r.GET(testTemplatePath, testTemplate)
	r.GET(testJSONPath, testJSON)
	r.GET(IndexViewPath, IndexView)
}

// 主页面
func IndexView(c echo.Context) error{
	return  c.String(200, "ok")
	
}

func testJSON(c echo.Context)error{
	return c.JSON(200, "test ok")
	 
}

func testTemplate(c echo.Context)error {

	return c.Render(200, "machine/test", echotool.H{
		"map": "测试页面",
	})

}
