package domain

type NodeReport struct {
	Requirements map[ResourceType]float64
}

func NewNodeReport(requirements map[ResourceType]float64) *NodeReport {
	r := new(NodeReport)

	r.Requirements = requirements

	return r
}
