package parser

import (
	gs "raw/src/global"
	"raw/src/structs"
	util "raw/src/utilities"
	"strings"
)

type ParseComputation struct {
	foundPrimaryKey bool
	foundUniqueKey  bool
}

func allowedTypes() []string {
	types := make([]string, 0)
	types = append(types, String)
	types = append(types, Integer)
	return types
}

func allowedAttributesInSchemaDefinition() []string {
	allowed := make([]string, 0)
	allowed = append(allowed, SchemaTypeName)
	allowed = append(allowed, PrimaryKey)
	allowed = append(allowed, UniqueKey)

	return allowed
}

func frameIncorrectTypeError(typeFound string) string {
	possibleTypes := strings.Join(allowedTypes(), ",")
	return "Incorrect type " + typeFound + " found when allowed types are " + possibleTypes
}

func frameTypeArgumentNotFound() string {
	return TypeArgumentNotFound
}

func frameIncorrectSchemaAttributeMesssage() string {
	allowed := strings.Join(allowedAttributesInSchemaDefinition(), ",")
	return "Schema attributes passed are incorrect. Allowed attributes are " + allowed
}

func frameMultiplePKMessage() string {
	return MultiplePKError
}

func checkForExistenceOfType(attributeSchema map[string]string) (gs.Error, bool) {
	typeVal, ok := attributeSchema["type"]
	error := gs.Error{Message: ""}
	success := true
	if ok {
		if !util.ContainsString(allowedTypes(), typeVal) {
			// Type not supported
			error.Message = frameIncorrectTypeError(typeVal)
			success = false
		}
	} else {
		// type argument not found
		error.Message = frameTypeArgumentNotFound()
		success = false
	}

	return error, success
}

func checkForValidSchemaAttributes(attributeSchema map[string]string) bool {
	return util.IsSubSet(allowedAttributesInSchemaDefinition(), util.GetAllKeys(attributeSchema))
}

func hasPrimaryKey(attributeSchema map[string]string) bool {
	_, ok := attributeSchema[PrimaryKey]
	if ok {
		return true
	}

	return false
}

// GetDefaults This function return the default values to be put
// in case of empty values specified with the schema
func GetDefaults(t string) string {
	switch t {
	case String:
		return EmptyString
	case Integer:
		return EmptyNone
	default:
		return EmptyNone
	}
}

// ExtractPrimaryKey Extracts the primary key from the schema. Returns an empty
// string and false if primary key is not found
func ExtractPrimaryKey(schema structs.Schema) (string, bool) {
	pKey := ""
	found := false
	for attr, attributeSchema := range schema {
		hPK := hasPrimaryKey(attributeSchema)
		if hPK {
			pKey = attr
			found = true
			break
		}
	}
	return pKey, found
}

func Parse(schema structs.Schema) (gs.Error, bool) {
	error := gs.Error{Message: ""}
	success := true
	computation := ParseComputation{
		foundPrimaryKey: false,
		foundUniqueKey:  false,
	}
	for _, attributeSchema := range schema {
		if checkForValidSchemaAttributes(attributeSchema) {
			tErr, tOk := checkForExistenceOfType(attributeSchema)
			if tOk {
				/**
				The following checks for existence of multiple primary keys
				inside the schema definition
				*/
				hPK := hasPrimaryKey(attributeSchema)
				if hPK && computation.foundPrimaryKey {
					error.Message = frameMultiplePKMessage()
					success = false
					break
				}

				if hPK {
					computation.foundPrimaryKey = true
				}
			} else {
				error.Message = tErr.Message
				success = false
				break
			}
		} else {
			error.Message = frameIncorrectSchemaAttributeMesssage()
			success = false
			break
		}
	}

	return error, success

}
