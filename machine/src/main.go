package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	//"os"
	_ "routes"
	"tool/gintool"
	"tool/inicfg"
)

/*
func init() {
	cfg, err := inicfg.Newcfg(os.Getenv("ETCDIR"))
	if err != nil {
		panic(err)

	}
	//test
	_, err = cfg.GetSection("master")
	if err != nil {
		panic(err)
	}
}
*/
func main() {

	r := gintool.Default()
	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.ReleaseMode)
	tpl := template.Must(template.New("machine").ParseGlob("./public/html/*.html"))
	gintool.SetHTMLTemplate(tpl)

	r.Static("/public/js", "./public/js")
	r.Static("/public/css", "./public/css")
	r.Static("/public/images", "./public/images")

	mapWeb, err := inicfg.Getcfg().GetSection("web/server/uri")
	if err != nil {
		panic(err)
	}

	r.Run(mapWeb["machine_server"])
}
