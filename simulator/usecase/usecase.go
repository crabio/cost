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
	nodesSimulations := make(map[string]*domain.NodeSimulation)

	// 1. Create empty report for each node
	for id := range sc.Nodes {
		nodesSimulations[id] = domain.NewNodeSimulation()
	}

	// 2. Search all clients
	clients, err := suc.ccuc.CreateNodesChains(sc)
	if err != nil {
		return nil, err
	}

	// 3. Requests flows gradient descent
	for _, n := range clients {
		suc.requestsFlowGradientDescent(nodesSimulations, n, nil)
	}

	// 4. Go through all nodes and calc consumption
	for _, ns := range nodesSimulations {
		for rf, as := range ns.RequestsFlows {
			for _, a := range as {
				for _, r := range a.Requirements {
					if nsr, ok := ns.Requirements[r.Resource]; ok {
						ns.Requirements[r.Resource] = r.Resource.Sum(nsr, r.Calc(rf))
					} else {
						ns.Requirements[r.Resource] = r.Calc(rf)
					}
				}
			}
		}
	}

	return domain.NewReport(), nil
}

func (suc *simulatorUsecase) requestsFlowGradientDescent(ns map[string]*domain.NodeSimulation, n *domain.Node, inrfs map[*domain.RequestsFlow][]*domain.Action) {
	// Add input requests flows
	ns[n.ID].RequestsFlows = inrfs

	// Create map with request flows for next nodes
	outrfs := make(map[*domain.RequestsFlow][]*domain.Action)

	for rf := range inrfs {
		for _, link := range n.Links {
			if outrf, ok := outrfs[rf]; ok {
				outrf = append(outrf, link.Action)
			} else {
				outrfs[rf] = []*domain.Action{link.Action}
			}
		}
	}

	// Go to children
	for _, link := range n.Links {
		suc.requestsFlowGradientDescent(ns, link.Child, outrfs)
	}
}
