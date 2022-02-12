package domain

type Model struct {
	ID               string
	Name             string
	Type             ModelType
	Params           []Model_Param
	AvailableActions map[string]Action
}

type ModelType string

const (
	ModelType_Custom ModelType = "custom"
	ModelType_Client ModelType = "client"
	ModelType_MQ     ModelType = "mq"
	ModelType_DB     ModelType = "db"
)

type Model_Param struct {
	Name          string        `yaml:"name"`
	UnitOfMeasure UnitOfMeasure `yaml:"uom"`
}
