package global

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
)

var GlobalCache cache.Cache

func init() {
	bm, err := cache.NewCache("memory", `{"interval":3600}`)
	if nil != err {
		logs.Error("init cache failed. cause: ", err)
		panic(err)
	}
	GlobalCache = bm
	logs.Debug("cache init.")
}

func Set(cache *cache.Cache) {
	GlobalCache = *cache
}
