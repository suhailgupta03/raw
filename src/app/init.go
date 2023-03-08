package app

import (
	"os"
	"raw/src/toml"
)

type Constants struct {
	AppName      string `json:"app_name"`
	DatabaseRoot string `json:"database_root"`
}

func InitConstants(config toml.RawConfig) Constants {
	constants := Constants{
		AppName:      config["app_name"].(string),
		DatabaseRoot: config["db_root"].(string),
	}
	return constants
}

func InitConfig() (toml.RawConfig, error) {
	configParser := toml.Parser()
	data, err := os.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}
	c, cErr := configParser.Unmarshal(data)
	if cErr != nil {
		panic(cErr)
	}

	return c, cErr
}
