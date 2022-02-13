package domain

type NodeSimulation struct {
	RequestsFlows map[*RequestsFlow][]*Action
	// Map with requirements
	// Key - resource type
	// Value - value
	Requirements map[ResourceType]float64
}

func NewNodeSimulation() *NodeSimulation {
	return new(NodeSimulation)
}
