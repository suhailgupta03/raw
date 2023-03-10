package api

import "raw/src/structs"

type CreateResponse struct {
	Data structs.CreateDataMap
	Err  error
}
