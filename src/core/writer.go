package core

import (
	"errors"
	"raw/src/parser"
	"raw/src/structs"
	"raw/src/utilities"
	"regexp"
	"strconv"
)

type Writer interface {
	Write()
}

type Document struct {
	root string
}

func InitDocument(root string) *Document {
	re := regexp.MustCompile("/$") // TODO: Use OS.Separator
	normRoot := re.ReplaceAllString(root, "")
	doc := Document{
		root: normRoot,
	}
	return &doc
}

func extractPKWithCorrectType(pkValue interface{}) string {
	pkType := utilities.TypeOf(pkValue)
	// Reference: https://pkg.go.dev/encoding/json#Unmarshal
	switch pkType {
	case "string":
		return pkValue.(string)
	case "int":
		return strconv.Itoa(pkValue.(int))
	case "float64":
		return strconv.FormatFloat(pkValue.(float64), 'E', -1, 64)
	default:
		panic(PrimaryKeyTypeExtractionError)
	}
}
func (d *Document) Write(schemaName string, schema structs.Schema, dataMap *structs.CreateDataMap) error {
	pk, fPK := parser.ExtractPrimaryKey(schema)
	dbUtil := DBUtil()
	pkValue := dbUtil.GenerateUUID()
	pkName := parser.DefaultPrimaryKey
	if fPK {
		pkName = pk
		// Since the values inside the data map could be of any type
		// it is important to detect the type and based on that
		// return the value after the type cast
		pkValue = extractPKWithCorrectType((*dataMap)[pkName])
	} else {
		// Update the data map with the system generated primary key
		(*dataMap)[pkName] = pkValue
	}

	// Generate the directory if it does not exist
	disk := InitDiskConfiguration(d.root)
	// Create the directory for primary keys
	// If the directory already exists, the directory
	// will not be recreated
	directory := PathJoiner(true, d.root, schemaName, dbUtil.GetPrimaryKeyRoot())
	_, dErr := disk.CreateDirectoryChain(directory)
	// Once the directory to store the primary keys has been created
	// we need to calculate the hash of the primary key. The hash
	// will be the name of the document that will be stored
	// inside the primary key directory
	if dErr != nil {
		panic(dErr)
	}
	primaryKeyHash := dbUtil.GeneratePrimaryKeyHash(pkValue)
	dataBytes := []byte(utilities.Stringify(*dataMap))
	recordPath := PathJoiner(false, directory, primaryKeyHash+disk.GetFileExtension())
	writeErr := disk.Write(recordPath, dataBytes)
	if writeErr != nil {
		if writeErr.Error() == RecordAlreadyExists {
			return errors.New(PrimaryKeyConstraintViolated + " (" + pkName + ")")
		}
		// Failed to create the record
		return writeErr
	}

	return nil
}
