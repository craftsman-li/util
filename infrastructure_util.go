package util

import "github.com/astaxie/beego/logs"

/**
 * 判断是否发生了错误
 */
func IsError(err error) bool {
	if nil != err {
		logs.Error("error ---> %v", err)
		return true
	}
	return false
}
