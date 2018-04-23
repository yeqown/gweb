package main

import (
	"gweb/models"

	"bufio"
	"encoding/json"
	"os"
)

type ServerConfig struct {
	Port    int    `json:"Port"`
	Logpath string `json:"Logpath"`
}

type Config struct {
	ServerC *ServerConfig       `json:"ServerConfig"`
	MysqlC  *models.MysqlConfig `json:"MysqlConfig"`
	RedisC  *models.RedisConfig `json:"RedisConfig"`
}

var _instance *Config

func GetInstance() *Config {
	return _instance
}

func LoadConfig(filepath string) error {
	fp, err := os.Open(filepath)
	defer fp.Close()
	if err != nil {
		return err
	}
	body, _ := bufio.NewReader(fp).ReadBytes(0)
	if err = json.Unmarshal(body, &_instance); err != nil {
		return err
	}
	return nil
}
