package domain_test

import (
	"testing"

	"github.com/iakrevetkho/cost/domain"
	"github.com/stretchr/testify/assert"
)

func TestDataFlowChainBlock_IsLinkExists(t *testing.T) {
	node3 := domain.NewNode("3", "3", domain.NodeType_Component, nil, nil)
	link2 := domain.NewLink("2", 1, node3, domain.LinkType_Local, nil)
	node2 := domain.NewNode("2", "2", domain.NodeType_Custom, nil, []*domain.Link{link2})
	link1 := domain.NewLink("1", 1, node2, domain.LinkType_Internet, nil)
	node1 := domain.NewNode("1", "1", domain.NodeType_Client, nil, []*domain.Link{link1})

	link4 := domain.NewLink("4", 1, node2, domain.LinkType_Internet, nil)

	chain1 := domain.NewDataFlowChainBlock(nil, node1, node2, link1)
	chain2 := domain.NewDataFlowChainBlock(chain1, node2, node3, link2)

	assert.True(t, chain2.IsLinkExists(link1))
	assert.True(t, chain2.IsLinkExists(link2))
	assert.False(t, chain2.IsLinkExists(link4))
}
