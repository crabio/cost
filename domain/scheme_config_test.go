package domain_test

import (
	"testing"

	"github.com/iakrevetkho/cost/domain"
	"github.com/iakrevetkho/cost/domain/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSchemeConfigFromYaml(t *testing.T) {
	sc, err := domain.NewSchemeConfigFromYamlBytes(testdata.SchemeCfg)
	require.NoError(t, err)
	assert.Len(t, sc.Nodes, 6)
	assert.Len(t, sc.Links, 11)
	assert.Len(t, sc.Models, 5)

	// Check client node
	assert.Equal(t, domain.NodeType_Client, sc.Nodes["1"].Type)
	assert.NotNil(t, sc.Nodes["1"].RequestsFlow)
	assert.Equal(t, uint(10), sc.Nodes["1"].RequestsFlow.Qty)
	assert.Equal(t, uint(1000), sc.Nodes["1"].RequestsFlow.PeriodMs)
	assert.Equal(t, uint(10000), sc.Nodes["1"].RequestsFlow.MsgSize)
}
