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
	assert.Len(t, sc.Links, 5)
	assert.Len(t, sc.Models, 3)
}
