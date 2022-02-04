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
	Links []Link
}
