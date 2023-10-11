package main

import (
	"oi/core"
	"oi/global"
	"oi/service"
)

func main() {
	// 加载环境变量
	core.InitCore()
	// 初始化日志
	global.Log = core.InitLogger()

	service.InitService()

}
