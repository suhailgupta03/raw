package api

import (
	"raw/src/core"
	"raw/src/parser"
	"raw/src/utilities"
)

type Schema map[string]map[string]string

type Model struct {
	Name   string `json:"name"`
	Schema Schema `json:"schema"`
}

type ModelBaseMethods interface {
	Create(create Create)
}

type ExtendedModelMethods interface {
	ModelBaseMethods
	MapParams(schemaKeys []string, queryKeys []string)
	ValidateQueryParameters()
}

func (m Model) Create(c Create) {
	if !utilities.IsEmptyString(m.Name) && !utilities.IsNil(m.Schema) {
		dataMap := m.MapParams(&c)
		document := core.Document{DataMap: dataMap}
		document.Write()
	} else {
		panic(SchemaNotInitialized)
	}
}

func (m Model) MapParams(c *Create) core.CreateDataMap {
	// Get all the keys defined in the schema
	schemaKeys := utilities.GetAllKeys(m.Schema)
	// Create a map that has keys derived from the above slice
	// Find the corresponding values in the 'values' interface
	// e.g. first_name in the schema should also exist in the
	// values interface and the value of the first_name
	// will be what is defined in the values interface
	// Note: if any key is not present in the value interface,
	// assign a default-value
	mappedParams := map[string]any{}
	for i := 0; i < len(schemaKeys); i++ {
		_, kExists := c.Values[schemaKeys[i]]
		if kExists {
			mappedParams[schemaKeys[i]] = c.Values[schemaKeys[i]]
		} else {
			attributeType := m.Schema[schemaKeys[i]][parser.SchemaTypeName]
			defaultValue := parser.GetDefaults(attributeType)
			mappedParams[schemaKeys[i]] = defaultValue
		}
	}
	// Once the map is prepared, this can now be used to
	// write to the database
	return mappedParams
}

// ValidateQueryParameters Verify if the parameters defined in the schema
// match those of the parameters passed in the database function
func (m Model) ValidateQueryParameters() {

}

func (m Model) Define() Model {
	// Verify if the schema passed is correct
	pErr, pOK := parser.Parse(m.Schema)
	if !pOK {
		// Invalid schema passed
		panic(pErr.Message)
	}

	return m
}
