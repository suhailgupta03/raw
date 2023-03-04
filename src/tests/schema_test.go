package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"raw/src/parser"
	"testing"
)

func TestParserWhenDataPassedIsCorrect(t *testing.T) {
	dataFileName := t.Name() + ".json"
	data, err := os.ReadFile(dataFileName)
	if err != nil {
		panic(err)
	}

	schema := map[string]map[string]string{}
	json.Unmarshal(data, &schema)
	parseError, parseSuccess := parser.Parse(schema)
	assert.True(t, parseSuccess, "Parse function should succeed")
	assert.Equal(t, "", parseError.Message)
}

func TestParserWhenMultiplePKPassed(t *testing.T) {
	dataFileName := t.Name() + ".json"
	data, err := os.ReadFile(dataFileName)
	if err != nil {
		panic(err)
	}
	schema := map[string]map[string]string{}
	json.Unmarshal(data, &schema)
	parseError, parseSuccess := parser.Parse(schema)
	assert.False(t, parseSuccess, "Parse function should fail")
	assert.Equal(t, parser.MultiplePKError, parseError.Message)
}

func TestParserWhenTypeMissing(t *testing.T) {
	dataFileName := t.Name() + ".json"
	data, err := os.ReadFile(dataFileName)
	if err != nil {
		panic(err)
	}
	schema := map[string]map[string]string{}
	json.Unmarshal(data, &schema)
	parseError, parseSuccess := parser.Parse(schema)
	assert.False(t, parseSuccess, "Parse function should fail")
	assert.Equal(t, parser.TypeArgumentNotFound, parseError.Message)
}

func TestParserWhenInvalidAttributesPassed(t *testing.T) {
	dataFileName := t.Name() + ".json"
	data, err := os.ReadFile(dataFileName)
	if err != nil {
		panic(err)
	}
	schema := map[string]map[string]string{}
	json.Unmarshal(data, &schema)
	parseError, parseSuccess := parser.Parse(schema)
	assert.False(t, parseSuccess, "Parse function should fail")
	assert.Contains(t, parseError.Message, "Schema attributes passed are incorrect")
}
