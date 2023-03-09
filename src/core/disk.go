package core

import (
	"errors"
	"os"
	"raw/src/compression"
	"regexp"
	"sync"
)

type Disk interface {
	CreateDirectory(folderName string) (bool, error)
	CreateDirectoryChain(chainPath string) (bool, error)
	Write(fileName string, data []byte) error
	Exists(fileName string) bool
}

type File interface {
	GetFileExtension() string
}

type DiskConfiguration struct {
	compression string
}

func InitDiskConfiguration(root string) *DiskConfiguration {
	err := os.Mkdir(root, DefaultDirectoryPermission)
	if err != nil && !os.IsExist(err) {
		// Creation of the root directory is a must before proceeding
		panic(err)
	}
	// Currently only gzip compression available
	return &DiskConfiguration{compression: GzipCompression}
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

func (dc *DiskConfiguration) GetFileExtension() string {
	switch dc.compression {
	case GzipCompression:
		return ".gz"
	default:
		return ".json"
	}
}

// Global lock to ensure that only one write operation runs at a time
var diskWriteLock sync.Mutex

// This method provides isolation to all the callers. It uses RWMutex
// to prevent data overwrite or inconsistent data
func (dc *DiskConfiguration) Write(fileName string, b []byte) error {
	diskWriteLock.Lock()
	defer diskWriteLock.Unlock()

	if dc.Exists(fileName) {
		return errors.New(RecordAlreadyExists)
	}

	if dc.compression == GzipCompression {
		gzip := new(compression.Gzip)
		b, _ = gzip.Compress(b)
	}
	err := os.WriteFile(fileName, b, DefaultFilePermission)
	return err
}

func (dc *DiskConfiguration) Exists(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
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
