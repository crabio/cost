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

func Test(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	ccuc := usecase.NewChainCreatorUsecase()

	sc, err := domain.NewSchemeConfigFromYamlBytes(testdata.SchemeCfg)
	require.NoError(t, err)

	chains, err := ccuc.CreateChains(sc)
	require.NoError(t, err)
	assert.Len(t, chains, 1)
}
