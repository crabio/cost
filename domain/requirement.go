package domain

import "github.com/sirupsen/logrus"

type Requirement struct {
	Resource ResourceType    `yaml:"resource"`
	Type     RequirementType `yaml:"type"`
	Value    float64         `yaml:"value"`
}

type RequirementType string

const (
	RequirementType_Once           RequirementType = "once"
	RequirementType_PerRequest     RequirementType = "per-request"
	RequirementType_PerRequestByte RequirementType = "per-request-byte"
)

func (r Requirement) Calc(rf *RequestsFlow) (float64, error) {
	if rf == nil {
		return 0, ErrNilRequirement
	}

	switch r.Type {
	case RequirementType_Once:
		return r.Value, nil

	case RequirementType_PerRequest:
		return r.Value * float64(rf.RequestsPerSecond()), nil

	case RequirementType_PerRequestByte:
		return r.Value * float64(rf.MsgSize), nil

	default:
		logrus.WithField("type", r.Type).Error(ErrUnknownRequirementType)
		return 0, ErrUnknownRequirementType
	}
}
