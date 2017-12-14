package gintool

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"sync"
	//"tool/inicfg"
)

//var temp *template.Template
var (
	// Global instance
	defaultE     *gin.Engine
	defaultELock = sync.Mutex{}
)

//设置路由，设置过滤器
func Default() *gin.Engine {
	defaultELock.Lock()
	defer defaultELock.Unlock()

	if defaultE == nil {
		engine := gin.Default()
		/*
		   if temp!=nil{
		       engine.SetHTMLTemplate(temp)
		   }
		*/
		//engine.Use(gin.Logger(), gin.Recovery())
		defaultE = engine

		//为啥这里的过滤器才生效？？？谁能告诉我！！！
/*
		// redis session

		mapEtc, err := inicfg.Getcfg().GetSection("redis")
		if err != nil {
			panic(err)
		}
		store_redis,err:= sessions.NewRedisStore(4, "tcp", mapEtc["web_redis_uri"], "",[]byte("machine_user_redis"))
		if err!=nil{
			panic(err)
		}
		defaultE.Use(sessions.Sessions("machine_user_redis", store_redis))
		*/

		store_cookie:=sessions.NewCookieStore([]byte("machine_user_cookie"))

		
		defaultE.Use(sessions.Sessions("machine_user_cookie", store_cookie))
		//defaultE.Use(auth.Auth)
		// cookie session
	}
	return defaultE
}

func SetHTMLTemplate(templ *template.Template) {
	defaultE.HTMLRender = render.HTMLProduction{Template: templ}
}
