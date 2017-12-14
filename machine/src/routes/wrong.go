package routes

//信息
const (
	//成功信息
	success="操作成功"
	//系统错误
	dbWrong    = "数据库异常"
	paramWrong = "请求参数异常"
	sysWrong   = "系统服务异常"
)

func paramWrongFormat(varname string) string {
	return paramWrong + ":" + varname
}
