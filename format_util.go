package util

import "github.com/craftsman-li/beego/orm"

func DateTimeFormat(in orm.DateTimeField) (out string) {
	return in.Value().Format(DateTimeFormatString)
}
