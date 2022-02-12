package domain

type NodeType string

const (
	NodeType_Client    = "client"
	NodeType_Custom    = "custom"
	NodeType_Component = "component"
)

type Node struct {
	ID    string
	Name  string
	Type  NodeType
	Model *Model
	Links []*Link
}

func NewNode(id string, name string, nodeType NodeType, model *Model, links []*Link) *Node {
	n := new(Node)

	n.ID = id
	n.Name = name
	n.Type = nodeType
	n.Model = model
	n.Links = links

	return n
}
