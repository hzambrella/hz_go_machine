package routes

import (
	"fmt"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

// 读取form值
func FormValue(c *gin.Context, key string) string {
	req := c.Request
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req.FormValue(key)
}

func DumpRequest(c *gin.Context) {
	data, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
