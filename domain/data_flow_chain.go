package domain

type DataFlowChain struct {
	Prev      *DataFlowChain
	StartNode *Node
	EndNode   *Node
	Link      *Link
}

func NewDataFlowChain(prev *DataFlowChain, startNode *Node, endNode *Node, link *Link) *DataFlowChain {
	c := new(DataFlowChain)

	c.Prev = prev
	c.StartNode = startNode
	c.EndNode = endNode
	c.Link = link

	return c
}
