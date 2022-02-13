package domain

import (
	"bytes"
)

type Action struct {
	Name         string               `yaml:"name"`
	Direction    Action_DirectionType `yaml:"direction"`
	Requirements []*Requirement       `yaml:"requirements"`
}

type Action_DirectionType string

const (
	Action_DirectionType_In  Action_DirectionType = "in"
	Action_DirectionType_Out Action_DirectionType = "out"
)

func (a *Action) String() string {
	var buf bytes.Buffer

	buf.WriteByte('{')
	buf.WriteString("name: ")
	buf.WriteString(a.Name)
	buf.WriteString(", direction: ")
	buf.WriteString(string(a.Direction))
	buf.WriteString(", requirements: [")

	for i, r := range a.Requirements {
		buf.WriteString(r.String())
		if i < (len(a.Requirements) - 1) {
			buf.WriteByte(',')
		}
	}
	buf.WriteString("]}")

	return buf.String()
}
