package domain

type ModelType string

const (
	ModelType_MQ = "mq"
	ModelType_DB = "db"
)

type Model struct {
	Name string
	Type ModelType
	// TODO Add actions
}
