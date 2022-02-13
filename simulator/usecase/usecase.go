package usecase

import (
	chain_creator_usecase "github.com/iakrevetkho/cost/chain_creator/usecase"
	"github.com/iakrevetkho/cost/domain"
)

type SimulatorUsecase interface {
	Simulate(sc *domain.SchemeConfig) (*domain.Report, error)
}

type simulatorUsecase struct {
	ccuc chain_creator_usecase.ChainCreatorUsecase
}

func NewSimulatorUsecase(ccuc chain_creator_usecase.ChainCreatorUsecase) SimulatorUsecase {
	suc := new(simulatorUsecase)
	suc.ccuc = ccuc
	return suc
}

func (suc *simulatorUsecase) Simulate(sc *domain.SchemeConfig) (*domain.Report, error) {
	r := domain.NewReport()

	// 1. Create empty report for each node
	for id := range sc.Nodes {
		r.NodeReports[id] = domain.NewNodeReport()
	}

	// 2. Search all clients
	clients, err := suc.ccuc.CreateNodesChains(sc)
	if err != nil {
		return nil, err
	}

	// 3. Requests flows gradient descent
	for _, n := range clients {
		suc.requestsFlowGradientDescent(r, n, nil)
	}

	return r, nil
}

func (suc *simulatorUsecase) requestsFlowGradientDescent(r *domain.Report, n *domain.Node, prfs []*domain.RequestsFlow) {
	// Add self requests flows
	if n.RequestsFlow != nil {
		r.NodeReports[n.ID].RequestsFlows = append(r.NodeReports[n.ID].RequestsFlows, n.RequestsFlow)
	}

	// Add parents requests flows
	r.NodeReports[n.ID].RequestsFlows = append(r.NodeReports[n.ID].RequestsFlows, prfs...)

	// Go to children
	for _, link := range n.Links {
		suc.requestsFlowGradientDescent(r, link.Child, r.NodeReports[n.ID].RequestsFlows)
	}
}
