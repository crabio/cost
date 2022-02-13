package domain

type NodeReport struct {
	RequestsFlow *RequestsFlow
	Requirements []*Requirement
}

func NewNodeReport() *NodeReport {
	return new(NodeReport)
}
