package main

import (

	"tool/echotool"
	
	"html/template"
	_ "routes"
)

const (
	testJSONPath     = "/machine/json"
	testTemplatePath = "/machine/template"
	IndexViewPath    = "/machine/testdb"
)

func main() {
	
	e := echotool.Default()

	
	//gin.SetMode(gin.ReleaseMode)
	tpl := template.Must(template.New("machine").ParseGlob("./public/html/*.html"))
	e.Renderer=&echotool.Template{tpl}
	e.Static("/public/js", "./public/js")
	e.Static("/public/css", "./public/css")
	e.Static("/public/images", "./public/images")

	//r.LoadHTMLGlob("public/html/*")

	e.Start(":8080")
}

/*
func IndexView(c *gin.Context) {
    c.String(200,"ok")
}

func testJSON(c *gin.Context) {
    c.JSON(200, "test ok")
}

func testTemplate(c *gin.Context) {
    c.HTML(200, "machine/test",gin.H{
		"map":"测试页面",
	})
}
*/
