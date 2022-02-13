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

	assert.Len(t, r.NodeReports["1"].Requirements, 0)

	assert.Len(t, r.NodeReports["2"].Requirements, 1)

	assert.Len(t, r.NodeReports["3"].Requirements, 3)
	// assert.Equal(t, map[domain.ResourceType]float64{
	// 	domain.ResourceType_Cpu:            100000,
	// 	domain.ResourceType_NetworkReceive: 10000,
	// 	domain.ResourceType_Ram:            10000,
	// }, r.NodeReports["3"].Requirements)

	assert.Len(t, r.NodeReports["4"].Requirements, 0)

	assert.Len(t, r.NodeReports["5"].Requirements, 3)

	assert.Len(t, r.NodeReports["6"].Requirements, 3)
}
