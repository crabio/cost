package domain_test

import (
	"testing"

	"github.com/iakrevetkho/cost/domain"
	"github.com/iakrevetkho/cost/domain/testdata"
	"github.com/sirupsen/logrus"
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

func TestSchemeConfigToScheme(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	sc, err := domain.NewSchemeConfigFromYamlBytes(testdata.SchemeCfg)
	require.NoError(t, err)

	s, err := sc.ToScheme()
	require.NoError(t, err)
	assert.Len(t, s.Roots, 2)

	assert.Equal(t, s.Roots[0].Name, "Client")
	assert.Equal(t, s.Roots[1].Name, "Gateway")
}
