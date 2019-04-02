package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kotoyuuko/bronya/logger"
)

type vhost struct {
	Name  []string
	Root  string
	Index []string
}

type config struct {
	Vhosts []vhost
}

// Config store parsed config
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
