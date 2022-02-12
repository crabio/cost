package domain

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type SchemeConfig struct {
	Nodes  map[string]*SchemeConfig_Node  `yaml:"nodes"`
	Links  map[string]*SchemeConfig_Link  `yaml:"links"`
	Models map[string]*SchemeConfig_Model `yaml:"models"`
}

type SchemeConfig_Node struct {
	Name string   `yaml:"name"`
	Type NodeType `yaml:"type"`
	// Model id
	Model  string             `yaml:"model"`
	Params map[string]float64 `yaml:"params"`
}

type SchemeConfig_Link struct {
	Seq uint `yaml:"seq"`
	// Start and End node id which link connects
	Start      string   `yaml:"start"`
	End        string   `yaml:"end"`
	Type       LinkType `yaml:"type"`
	ActionName string   `yaml:"action"`
}

type SchemeConfig_Model struct {
	Name             string            `yaml:"name"`
	Type             ModelType         `yaml:"type"`
	Params           []Model_Param     `yaml:"param"`
	AvailableActions map[string]Action `yaml:"available-actions"`
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
