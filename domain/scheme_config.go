package domain

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type SchemeConfig struct {
	Nodes  map[string]*SchemeConfig_Node `yaml:"nodes"`
	Links  map[string]*SchemeConfig_Link `yaml:"links"`
	Models map[string]*Model             `yaml:"models"`
}

type SchemeConfig_Node struct {
	Name string   `yaml:"name"`
	Type NodeType `yaml:"type"`
	// Model id
	Model string `yaml:"model"`
}

type SchemeConfig_Link struct {
	Seq uint `yaml:"seq"`
	// Start and End node id which link connects
	Start string   `yaml:"start"`
	End   string   `yaml:"end"`
	Type  LinkType `yaml:"type"`
}

func NewSchemeConfigFromYaml(filePath string) (*SchemeConfig, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return NewSchemeConfigFromYamlBytes(fileBytes)
}

func NewSchemeConfigFromYamlBytes(in []byte) (*SchemeConfig, error) {
	c := new(SchemeConfig)

	if err := yaml.Unmarshal(in, c); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *SchemeConfig) ToScheme() (*Scheme, error) {
	s := new(Scheme)

	// Create map with all nodes
	// Key - ID
	// Value - Node
	nodesMap := make(map[string]*Node)
	// Create map with root nodes
	// Key - ID
	// Value - struct
	rootNodesMap := make(map[string]struct{})
	for id, node := range c.Nodes {
		logrus.WithField("node", node).Debug("node")
		nodesMap[id] = &Node{
			ID:   id,
			Name: node.Name,
			Type: node.Type,
		}
		rootNodesMap[id] = struct{}{}
	}

	for nodeId, node := range nodesMap {
		logrus.WithFields(logrus.Fields{"id": nodeId, "node": node}).Debug("node")
		for linkId, link := range c.Links {
			logrus.WithFields(logrus.Fields{"id": linkId, "link": link}).Debug("link")
			if link.Parent == nodeId {
				childNode, ok := nodesMap[link.Child]
				if !ok {
					logrus.WithField("link", link).Warn("unknown child node id")
				}

				node.Links = append(node.Links, Link{
					ID:    linkId,
					Seq:   link.Seq,
					Child: childNode,
					Type:  link.Type,
				})
			} else if link.Child == nodeId {
				// Delte node from root nodes
				delete(rootNodesMap, nodeId)
			}
		}
	}

	// Search root nodes
	for id, _ := range rootNodesMap {
		node := nodesMap[id]
		s.Roots = append(s.Roots, node)
	}

	return s, nil
}
