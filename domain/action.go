package domain

type Action struct {
	Name         string               `yaml:"name"`
	Direction    Action_DirectionType `yaml:"type"`
	Requirements []Requirement        `yaml:"requirements"`
}

type Action_DirectionType string

const (
	Action_DirectionType_In  Action_DirectionType = "in"
	Action_DirectionType_Out Action_DirectionType = "out"
)
