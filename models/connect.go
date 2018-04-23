package models

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	. "gweb/logger"
	"gweb/utils"
)

var (
	mysqlIns    *gorm.DB
	postgresIns *gorm.DB
	redisIns    *redis.Client
)

type MysqlConfig struct {
	Addr      string `json:"Addr"`
	Loc       string `json:"Loc"`
	Charset   string `json:"Charset"`
	Pool      int    `json:"pool"`
	ParseTime string `json:"ParseTime"`
}

type MongoConfig struct {
}

type RedisConfig struct {
	Addr        string `json:"Addr"`
	DB          int    `json:"DB"`
	Password    string `json:"Password"`
	PoolSize    int    `json:"PoolSize"`
	Timeout     int    `json:"Timeout"`
	MaxActive   int    `json:"MaxActive"`
	MaxIdle     int    `json:"MaxIdle"`
	IdleTimeout int    `json:"IdleTimeout"`
	Wait        bool   `json:"Wait"`
}

type PostgresConfig struct {
}

/*
 * connect to mysql
 */
func ConnectMysql(myc *MysqlConfig) {
	conStr := utils.Fstring("%sloc=%s&parseTime=%s&charset=%s",
		myc.Addr,
		myc.Loc,
		myc.ParseTime,
		myc.Charset,
	)
	AppL.Infof("Connect to mysql with Connect string: %s\n", conStr)
	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		AppL.Fatalf("Open mysql failed: %s\n", err.Error())
	}

	db.DB().SetMaxOpenConns(myc.Pool)
	db.DB().SetMaxIdleConns(myc.Pool / 2)
	db.LogMode(false)
	db.SingularTable(true)

	if err = db.DB().Ping(); err != nil {
		AppL.Fatalf("Ping mysql failed: %s", err.Error())
	}
	mysqlIns = db
}

// TODO: Connect to Mongo
func ConnectMongo(moc *MongoConfig) {

}

func ConnectRedis(rec *RedisConfig) {
	AppL.Info("Connect to redis")
	redisIns = redis.NewClient(&redis.Options{
		Addr:     rec.Addr,
		Password: rec.Password,
		DB:       rec.DB,
		PoolSize: rec.PoolSize,
	})
}

// TODO: Connect postgres
func ConnectPostgres(poc *PostgresConfig) {

}

/*
 * Get db connection
 * Mysql / Postgres / Redis / Mongo .etc
 */
func getMysqlDB() *gorm.DB {
	return mysqlIns
}

func getPostgresDB() *gorm.DB {
	return postgresIns
}

func getRedisDB() *redis.Client {
	return redisIns
}

// TODO: get mongo connection
func GetMongoDB() {

}
