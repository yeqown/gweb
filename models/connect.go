package models

import (
	"github.com/jinzhu/gorm"
)

var (
	mysqlIns    *gorm.DB
	postgresIns *gorm.DB
)

type MysqlConfig struct {
}

type MongoConfig struct {
}

type RedisConfig struct {
}

type PostgresConfig struct {
}

/*
 * connect to mysql
 */
func ConnectMysql(myc *MysqlConfig) {

}

func ConnectMongo(moc *MongoConfig) {

}

func ConnectRedis(rec *RedisConfig) {

}

func ConnectPostgres(poc *PostgresConfig) {

}

/*
 * get db connection
 */
func GetMysqlDB() *gorm.DB {
	return mysqlIns
}

func GetPostgresDB() {
	return postgresIns
}

func GetRedisDB() {

}
func GetMongoDB() {

}
