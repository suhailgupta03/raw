package src

import (
	"fmt"
	p "raw/src/parser"
)

func Tester() {
	userSchema := make(map[string]map[string]string)
	firstNameSchema := make(map[string]string)
	firstNameSchema["type"] = "string"

	secondNameSchema := make(map[string]string)
	secondNameSchema["type"] = "string"

	ageSchema := make(map[string]string)
	ageSchema["type"] = "integer"

	rollNumberSchema := make(map[string]string)
	rollNumberSchema["type"] = "integer"
	rollNumberSchema["primary"] = "true"

	userSchema["first_name"] = firstNameSchema
	userSchema["second_name"] = secondNameSchema
	userSchema["age"] = ageSchema
	userSchema["roll_number"] = rollNumberSchema

	parseError, parseSuccess := p.Parse(userSchema)
	if !parseSuccess {
		fmt.Println(parseError.Message)
	} else {
		fmt.Println("Schema parse success")
	}
}
