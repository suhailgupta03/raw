package tests

import (
	"os"
	"raw/src/app"
	"raw/src/core"
	"regexp"
	"strings"
	"testing"
)

type TestConfig struct {
	DatabaseRoot string `json:"db_root"`
}

func prepareConfKey(key string) string {
	return key + "_test"
}
func InitTestConfig() TestConfig {

	conf, confErr := app.InitConfig()
	if confErr != nil {
		panic(confErr)
	}

	testConfig := TestConfig{
		DatabaseRoot: conf[prepareConfKey("db_root")].(string),
	}

	return testConfig
}
func GetTestFileData(t *testing.T, filename string) []byte {
	reg := regexp.MustCompile(".go$")
	filename = reg.ReplaceAllString(filename, "")
	filePathSplit := strings.Split(filename, string(os.PathSeparator))
	filename = filePathSplit[len(filePathSplit)-1]
	dataFileName := core.PathJoiner(false, "data", filename, t.Name()+".json")
	data, err := os.ReadFile(dataFileName)
	if err != nil {
		panic(err)
	}
	return data
}
