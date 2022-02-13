package domain

type NodeReport struct {
	Requirements []*Requirement
}

func NewNodeReport(requirements []*Requirement) *NodeReport {
	r := new(NodeReport)

	r.Requirements = requirements

	return r
}
