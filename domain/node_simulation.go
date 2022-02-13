package domain

type NodeSimulation struct {
	RequestsFlows map[*RequestsFlow][]*Action
	Requirements  []*Requirement
}

func NewNodeSimulation() *NodeSimulation {
	return new(NodeSimulation)
}
