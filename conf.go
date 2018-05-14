package gweb

import (
	"bufio"
	"encoding/json"
	"os"
)

type ServerConfig struct {
	Port    int    `json:"Port"`
	Logpath string `json:"Logpath"`
}

type RpcServerConfig struct {
	Host      string `json:"host"`       // host
	Port      int    `json:"port"`       // port
	Path      string `json:"path"`       // path
	DebugPath string `json:"debug_path"` // debug_path
	Network   string `json:"network"`    // tcp http
}

type Config struct {
	ServerC *ServerConfig    `json:"ServerConfig"`
	RpcC    *RpcServerConfig `json:"RpcServerConfig"`
}

var _instance *Config

// GetInstance get global config instance
func GetInstance() *Config {
	return _instance
}

// LoadConfig read config.json file and pares to
// Golang Struct config instance varible
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
