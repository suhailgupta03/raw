package core

type CreateDataMap map[string]any

type Writer interface {
	Write()
}
type Document struct {
	DataMap CreateDataMap `json:"data_map"`
}

func (d Document) Write() {

}
