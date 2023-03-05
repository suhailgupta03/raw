package src

import (
	"raw/src/api"
	"raw/src/app"
)

func Tester(constants app.Constants) {
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

	model := api.Model{
		Name:   "User",
		Schema: userSchema,
	}

	model = model.Define()
	createSchema := api.Create{
		Values: map[string]any{
			"first_name":  "suhail",
			"second_name": "gupta",
			"roll_number": 44,
		},
	}

	model.Create(createSchema)

}
