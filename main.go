package main

import (
	"Gin_Start/bootstrap"
	"Gin_Start/global"
)

func main() {
	bootstrap.InitializeConfig()

	//初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("日志初始化成功")

	//初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	global.App.Log.Info("数据库初始化成功")

	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()
	bootstrap.InitializeValidator() //初始化验证器
	bootstrap.RunServer()
}
