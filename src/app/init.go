package app

import (
	"os"
	"raw/src/core"
	"raw/src/toml"
)

type Constants struct {
	AppName      string `json:"app_name"`
	DatabaseRoot string `json:"database_root"`
}

func InitConstants(config toml.RawConfig) Constants {
	homeDirName, hErr := os.UserHomeDir()
	if hErr != nil {
		panic(hErr.Error())
	}

	constants := Constants{
		AppName:      config["app_name"].(string),
		DatabaseRoot: core.PathJoiner(false, homeDirName, config["db_root"].(string)),
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
