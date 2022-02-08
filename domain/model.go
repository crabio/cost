package domain

type ModelType uint

const (
	ModelType_NA ModelType = iota
	ModelType_Custom
	ModelType_Client
	ModelType_MQ
	ModelType_DB
)

type Model struct {
	ID               string
	Name             string
	Type             ModelType
	AvailableActions []Action
}

type Action struct {
	Name         string
	Direction    Action_DirectionType
	Requirements []Requirement
}

type Action_DirectionType uint

const (
	Action_DirectionType_In Action_DirectionType = iota
	Action_DirectionType_Out
)

type Requirement struct {
	ResourceType RequirementResourceType
	Type         RequirementType
}

type RequirementResourceType uint

const (
	RequirementResourceType_CPU RequirementResourceType = iota
	RequirementResourceType_RAM
)

type RequirementType uint

const (
	RequirementType_Once RequirementType = iota
	RequirementType_PerRequest
)
