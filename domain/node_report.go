package domain

type NodeReport struct {
	RequestsFlows []*RequestsFlow
	Requirements  []*Requirement
}

func NewNodeReport() *NodeReport {
	return new(NodeReport)
}
