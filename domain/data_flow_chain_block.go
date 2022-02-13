package domain

type DataFlowChainBlock struct {
	Prev      *DataFlowChainBlock
	StartNode *Node
	EndNode   *Node
	Link      *Link
}

func NewDataFlowChainBlock(prev *DataFlowChainBlock, startNode *Node, endNode *Node, link *Link) *DataFlowChainBlock {
	c := new(DataFlowChainBlock)

	c.Prev = prev
	c.StartNode = startNode
	c.EndNode = endNode
	c.Link = link

	return c
}

func (b *DataFlowChainBlock) IsLinkExists(l *Link) bool {
	if b.Link == l {
		return true
	}

	// Check prev block
	if b.Prev != nil {
		return b.Prev.IsLinkExists(l)
	}

	return false
}
