package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"raw/src/parser"
	"raw/src/structs"
	"runtime"
	"testing"
)

func getDataForSchemaTest(t *testing.T) structs.Schema {
	_, filename, _, _ := runtime.Caller(0)
	data := GetTestFileData(t, filename)
	schema := structs.Schema{}
	json.Unmarshal(data, &schema)
	return schema
}
func TestParserWhenDataPassedIsCorrect(t *testing.T) {
	schema := getDataForSchemaTest(t)
	parseError, parseSuccess := parser.Parse(schema)
	assert.True(t, parseSuccess, "Parse function should succeed")
	assert.Equal(t, "", parseError.Message)
}

func TestParserWhenMultiplePKPassed(t *testing.T) {
	schema := getDataForSchemaTest(t)
	parseError, parseSuccess := parser.Parse(schema)
	assert.False(t, parseSuccess, "Parse function should fail")
	assert.Equal(t, parser.MultiplePKError, parseError.Message)
}

func TestParserWhenTypeMissing(t *testing.T) {
	schema := getDataForSchemaTest(t)
	parseError, parseSuccess := parser.Parse(schema)
	assert.False(t, parseSuccess, "Parse function should fail")
	assert.Equal(t, parser.TypeArgumentNotFound, parseError.Message)
}

func TestParserWhenInvalidAttributesPassed(t *testing.T) {
	schema := getDataForSchemaTest(t)
	parseError, parseSuccess := parser.Parse(schema)
	assert.False(t, parseSuccess, "Parse function should fail")
	assert.Contains(t, parseError.Message, "Schema attributes passed are incorrect")
}
