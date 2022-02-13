package domain

type LinkType string

const (
	LinkType_Internet = "internet"
	LinkType_Local    = "local"
)

type Link struct {
	ID     string
	Seq    uint
	Child  *Node
	Type   LinkType `yaml:"type"`
	Action *Action
}

func NewLink(id string, seq uint, child *Node, linkType LinkType, action *Action) *Link {
	l := new(Link)

	l.ID = id
	l.Seq = seq
	l.Child = child
	l.Type = linkType
	l.Action = action

	return l
}
