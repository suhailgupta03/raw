package main

import (
	t "raw/src"
	"raw/src/app"
	"raw/src/toml"
)

var (
	rawConfig    toml.RawConfig
	appConstants app.Constants
)

func main() {
	c, confErr := app.InitConfig()
	if confErr != nil {
		panic(confErr)
	}

	rawConfig = c
	appConstants = app.InitConstants(rawConfig)
	test()
}

func test() {
	t.Tester(appConstants)
}
