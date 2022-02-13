package usecase_test

import (
	"testing"

	"github.com/iakrevetkho/cost/chain_creator/usecase"
	"github.com/iakrevetkho/cost/domain"
	"github.com/iakrevetkho/cost/domain/testdata"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChainCreatorUsecase_CreateNodesChains(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	ccuc := usecase.NewChainCreatorUsecase()

	sc, err := domain.NewSchemeConfigFromYamlBytes(testdata.SchemeCfg)
	require.NoError(t, err)

	chains, err := ccuc.CreateNodesChains(sc)
	require.NoError(t, err)
	assert.Len(t, chains, 1)

	assert.Equal(t, "1", chains[0].ID)
	assert.Len(t, chains[0].Links, 1)
	assert.Equal(t, "2", chains[0].Links[0].Child.ID)

	child2 := chains[0].Links[0].Child
	assert.Len(t, child2.Links, 4)

	var child3, child5 *domain.Node
	for _, link := range child2.Links {
		if link.Child.ID == "3" {
			child3 = link.Child
		} else if link.Child.ID == "5" {
			child5 = link.Child
		}
	}

	require.NotNil(t, child3)
	assert.Len(t, child3.Links, 1)
	assert.Equal(t, "2", child3.Links[0].Child.ID)

	require.NotNil(t, child5)
	assert.Len(t, child5.Links, 1)
	assert.Equal(t, "4", child5.Links[0].Child.ID)

	child4 := child5.Links[0].Child
	assert.Len(t, child4.Links, 4)

	var child6 *domain.Node
	for _, link := range child4.Links {
		if link.Child.ID == "6" {
			child6 = link.Child
		}
	}
	require.NotNil(t, child6)
}

func TestChainCreatorUsecase_CreateChains(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	ccuc := usecase.NewChainCreatorUsecase()

	sc, err := domain.NewSchemeConfigFromYamlBytes(testdata.SchemeCfg)
	require.NoError(t, err)

	chains, err := ccuc.CreateChains(sc)
	require.NoError(t, err)
	assert.Len(t, chains, 1)
}
