package auth

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"fmt"
	//"tool/redis"
)

//鉴权拦截器
var UserSessNotFound = errors.New("user not found")

//用户信息
type UserSession struct {
	//uid_auth
	//AuthType string `json:"auth_type"`
	Uid int64 `json:"uid"`
}

const (
	UserSessionKey = "machine_user"
	ReqPathSessionKey="machine_req_path"
	loginViewPath  = "/machine/login/view"
)

//鉴权
//拦截器千万别用r.Use()，放到handler的第一行。
func Auth(c *gin.Context)bool {
	/*
	needAuth, err := NeedInterceptor(c.Request.URL.Path)
	if err != nil {
		c.AbortWithError(500, err)
		c.String(500, err.Error())
		return
	}
	if !needAuth {
		c.Next()
		return
	}
*/
	
	userSess, err := GetUserSession(c)
	if userSess!=nil{
		fmt.Println("auth:",userSess.Uid)
	}else{
		fmt.Println("马丹")
	}

	if err != nil {
		if err != UserSessNotFound {
			c.AbortWithError(500, err)
			c.String(500, err.Error())
			return false
		}else{
			if err:=SaveUserReqPath(c.Request.URL.Path,c);err!=nil{
				c.AbortWithError(500, err)
				c.String(500, err.Error())
				return false
			}
			c.Redirect(302, "/machine/login/view")
			return false
		}
	}

	if err := setUserSession(userSess, c); err != nil {
		c.AbortWithError(500, err)
		c.String(500, err.Error())
		return false
	}

	return true
}

/*
func isLogin(c *gin.Context)bool{
	)
	JSONsess.Get("user_login")
}
*/
// 设置会话
func SetUserSession(uid int64, c *gin.Context) error {
	uss := &UserSession{Uid: uid}
	return setUserSession(uss, c)
}

func setUserSession(user *UserSession, c *gin.Context) error {
	sess := sessions.Default(c)
	userAuth, err := json.Marshal(user)
	if err != nil {
		return err
	}
	sess.Options(sessions.Options{MaxAge: 7 * 24 * 60 * 60, Path: "/"})
	sess.Set(UserSessionKey, userAuth)
	if err:=sess.Save();err!=nil{
		return err
	}
	return nil
}

// 获取会话
func GetUserSession(c *gin.Context) (*UserSession, error) {
	sess := sessions.Default(c)
	uSessByte := sess.Get(UserSessionKey)
	if uSessByte == nil {
		return nil, UserSessNotFound
	}
	userSess := &UserSession{}
	if err := json.Unmarshal(uSessByte.([]byte), userSess); err != nil {
		return nil, err
	}
	return userSess, nil
}
/*
// 请求是否需要拦截鉴权，在$ETCDIR/intercepotor.ini设置不需拦截的。
func NeedInterceptor(paths string) (bool, error) {
	_, err := os.Stat(os.Getenv("ETCDIR") + "/interceptor.ini")
	if err != nil {
		if os.IsNotExist(err) {
			return true, errors.New(os.Getenv("ETCDIR") + "/interceptor.ini not found")
		}
	}

	mapExclude, err := inicfg.Getcfg().GetSection("intercetor_exclude")
	if err != nil {
		return true, err
	}

	exclude, ok := mapExclude["exclude_path"]
	if !ok {
		return true, errors.New("wrong config about interceptor,please check in " + os.Getenv("ETCDIR") + "/intercepotor.ini")
	}

	result := strings.Index(exclude, paths)
	if result < 0 {
		return true, nil
	}

	return false, nil
}
*/
//用户第一次请求的链接
func ClearUserReqPath(c *gin.Context){
	sess := sessions.Default(c)
	sess.Delete(ReqPathSessionKey)
}

func GetUserReqPath(c *gin.Context)string{
	sess := sessions.Default(c)
	return sess.Get(ReqPathSessionKey).(string)

}

func SaveUserReqPath(path string, c *gin.Context) error {
	sess := sessions.Default(c)
	sess.Options(sessions.Options{MaxAge: 7 * 24 * 60 * 60, Path: "/"})
	sess.Set(ReqPathSessionKey, path)
	if err:=sess.Save();err!=nil{
		return err
	}
	return nil
}

func (u *UserSession) GetUid() int64 {
	return u.Uid
}
