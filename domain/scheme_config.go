package domain

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type SchemeConfig struct {
	Nodes  map[uint]SchemeConfig_Node `yaml:"nodes"`
	Links  map[uint]SchemeConfig_Link `yaml:"links"`
	Models map[uint]Model             `yaml:"models"`
}

type SchemeConfig_Node struct {
	Name string   `yaml:"name"`
	Type NodeType `yaml:"type"`
	// Model id
	Model uint `yaml:"model"`
}

type SchemeConfig_Link struct {
	Seq uint `yaml:"seq"`
	// Parent node id
	Parent uint `yaml:"parent"`
	// Child node id
	Child uint     `yaml:"child"`
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
