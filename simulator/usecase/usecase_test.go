package usecase_test

import (
	"testing"

	chain_creator "github.com/iakrevetkho/cost/chain_creator/usecase"
	"github.com/iakrevetkho/cost/domain"
	"github.com/iakrevetkho/cost/domain/testdata"
	simulator "github.com/iakrevetkho/cost/simulator/usecase"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimulatorUsecase_Simulate(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	ccuc := chain_creator.NewChainCreatorUsecase()

	suc := simulator.NewSimulatorUsecase(ccuc)

	sc, err := domain.NewSchemeConfigFromYamlBytes(testdata.SchemeCfg)
	require.NoError(t, err)

	r, err := suc.Simulate(sc)
	require.NoError(t, err)
	assert.Len(t, r.NodeReports, 6)
}
