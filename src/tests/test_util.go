package tests

import (
	"os"
	"raw/src/app"
	"raw/src/core"
	"regexp"
	"strings"
	"testing"
)

func InitTestConfig() app.Constants {

	conf, confErr := app.InitConfig()
	if confErr != nil {
		panic(confErr)
	}

	return app.InitConstants(conf)

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
