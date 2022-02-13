package domain

type Requirement struct {
	Resource      ResourceType    `yaml:"resource"`
	Type          RequirementType `yaml:"type"`
	UnitOfMeasure UnitOfMeasure   `yaml:"uom"`
	Value         float64         `yaml:"value"`
}

type RequirementType string

const (
	RequirementType_Once           RequirementType = "once"
	RequirementType_PerRequest     RequirementType = "per-request"
	RequirementType_PerRequestByte RequirementType = "per-request-byte"
)
