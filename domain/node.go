package domain

import "sort"

type NodeType string

const (
	NodeType_Client    NodeType = "client"
	NodeType_Custom    NodeType = "custom"
	NodeType_Component NodeType = "component"
)

type Node struct {
	ID    string
	Name  string
	Type  NodeType
	Model *Model
	Links []*Link

	// Specified only for client node
	RequestsFlow *RequestsFlow
}

func NewNode(id string, name string, nodeType NodeType, model *Model, links []*Link) *Node {
	n := new(Node)

	n.ID = id
	n.Name = name
	n.Type = nodeType
	n.Model = model
	n.Links = links

	// Sort links by seq
	sort.Slice(n.Links, func(i, j int) bool {
		return n.Links[i].Seq < n.Links[j].Seq
	})

	return n
}

func NewClientNode(id string, name string, nodeType NodeType, model *Model, links []*Link, requestsFlow *RequestsFlow) *Node {
	n := NewNode(id, name, nodeType, model, links)

	n.RequestsFlow = requestsFlow

	return n
}
