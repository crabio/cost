package domain

type NodeSimulation struct {
	// Requests flows
	RequestsFlows []*RequestsFlow
	// Map with actions on node
	ActionsFlows map[*RequestsFlow][]*Action
	// Map with requirements
	// Key - resource type
	// Value - value
	Requirements map[ResourceType]float64
}

func NewNodeSimulation() *NodeSimulation {
	ns := new(NodeSimulation)

	ns.ActionsFlows = make(map[*RequestsFlow][]*Action)
	ns.Requirements = make(map[ResourceType]float64)

	return ns
}
