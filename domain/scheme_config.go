package domain

type SchemeConfig struct {
	Nodes  map[uint]SchemeConfig_Node
	Links  map[uint]SchemeConfig_Link
	Models map[uint]Model
}

type SchemeConfig_Node struct {
	Name string
	Type NodeType
	// Model id
	Model uint
}

type SchemeConfig_Link struct {
	Seq uint
	// Parent node id
	Parent uint
	// Child node id
	Child uint
	Type  LinkType
}
