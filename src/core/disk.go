package core

import (
	"os"
	"regexp"
)

type Disk interface {
	CreateDirectory(folderName string) (bool, error)
	CreateDirectoryChain(chainPath string) (bool, error)
	Write(fileName string, data []byte) (bool, error)
}

type DiskConfiguration struct {
}

func InitDiskConfiguration(root string) *DiskConfiguration {
	err := os.Mkdir(root, DefaultDirectoryPermission)
	if err != nil && !os.IsExist(err) {
		// Creation of the root directory is a must before proceeding
		panic(err)
	}
	return &DiskConfiguration{}
}

func (dc *DiskConfiguration) CreateDirectory(folderName string) (bool, error) {
	err := os.Mkdir(folderName, DefaultDirectoryPermission)
	if err != nil && !os.IsExist(err) {
		return false, err
	}

	return true, nil
}

func (dc *DiskConfiguration) CreateDirectoryChain(chainPath string) (bool, error) {
	err := os.MkdirAll(chainPath, DefaultDirectoryPermission)
	if err != nil && !os.IsExist(err) {
		return false, err
	}

	return true, nil
}

func (dc *DiskConfiguration) Write(fileName string, b []byte) error {
	err := os.WriteFile(fileName, b, DefaultFilePermission)
	return err
}

// PathJoiner This path joins the directory names with the os specific
// directory separator. If the path needs to end with the filename
// directoryPath must be set as false
func PathJoiner(onlyDirectory bool, pathNames ...string) string {
	joinedPath := ""
	re := regexp.MustCompile("/$") // TODO: Use OS.Separator
	for index, path := range pathNames {
		transformedPath := re.ReplaceAllString(path, "")
		if index == len(pathNames)-1 && !onlyDirectory {
			joinedPath += transformedPath
		} else {
			joinedPath += transformedPath + string(os.PathSeparator)
		}
	}

	return joinedPath
}
