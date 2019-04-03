package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/kotoyuuko/bronya/logger"
)

type fastcgi struct {
	Network string
	Address string
}

// Vhost 存储虚拟主机信息
type Vhost struct {
	Name    []string
	Root    string
	Index   []string
	Fastcgi fastcgi
}

type config struct {
	Listen  string
	Port    string
	Vhosts  []Vhost
	Default Vhost
}

// Config 存储从配置文件中读取并解析后的配置
var Config *config

func init() {
	logger.Info.Println("Parsing config file...")

	configJSON, err := ioutil.ReadFile("config.json")
	if err != nil {
		logger.Error.Fatalln(err)
	}

	Config = &config{}
	err = json.Unmarshal(configJSON, Config)
	if err != nil {
		logger.Error.Fatalln(err)
	}

	logger.Info.Println("Config file parsed.")
}

// SearchVhost 按照指定的域名查找虚拟主机
func SearchVhost(searchName string) (*Vhost, error) {
	for _, host := range Config.Vhosts {
		for _, name := range host.Name {
			if name == searchName {
				return &host, nil
			}
		}
	}

	return &Config.Default, errors.New("Vhost not found")
}
