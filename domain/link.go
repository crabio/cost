package domain

type LinkType string

const (
	LinkType_Internet = "internet"
	LinkType_Local    = "local"
)

type Link struct {
	ID    string
	Seq   uint
	Child *Node
	Type  LinkType `yaml:"type"`
}
