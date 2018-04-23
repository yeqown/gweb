package main

import (
	. "gweb/logger"
	"gweb/models"
)

/*
 * connect entry
 * like:
 */
func ConnectDBs() {
	// _ins := GetInstance()
	if _instance == nil {
		AppL.Fatal("Config instance is nil")
	}

	models.ConnectMysql(_instance.MysqlC)
	// models.ConnectMongo()
	models.ConnectRedis(_instance.RedisC)
	// models.ConnectPostgres()
}
