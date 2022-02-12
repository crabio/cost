package domain

type DataFlowChain struct {
	Prev      *DataFlowChain
	StartNode *Node
	EndNode   *Node
	Link      *Link
}
