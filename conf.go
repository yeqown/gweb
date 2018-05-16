package gweb

import (
// "bufio"
// "encoding/json"
// "os"
)

type ServerConfig struct {
	Port    int    `json:"Port"`    // HttpServer Port
	Logpath string `json:"Logpath"` // HttpServer Logpath
}

type RpcServerConfig struct {
	Host      string `json:"host"`       // RpcServer Host
	Port      int    `json:"port"`       // RpcServer Port
	Path      string `json:"path"`       // RpcSerevr Path
	DebugPath string `json:"debug_path"` // RpcServer Debug path
	Network   string `json:"network"`    // RpcServer Newwork
}

type MysqlConfig struct {
	Addr      string `json:"Addr"`
	Loc       string `json:"Loc"`
	Charset   string `json:"Charset"`
	Pool      int    `json:"pool"`
	ParseTime string `json:"ParseTime"`
}

type MongoConfig struct {
	Addrs     string `json:"Addrs"`
	Timeout   int64  `json:"Timeout"`
	PoolLimit int    `json:"PoolLimit"`
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
