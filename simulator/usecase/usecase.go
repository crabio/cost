package usecase

import (
	chain_creator_usecase "github.com/iakrevetkho/cost/chain_creator/usecase"
	"github.com/iakrevetkho/cost/domain"
	"github.com/sirupsen/logrus"
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

	// Links set required to prevent infinite loops
	linksSet := make(map[*domain.Link]struct{})

	for _, n := range clients {
		suc.requestsFlowGradientDescent(nodesSimulations, linksSet, n, nil)
	}

	// 4. Go through all nodes and calc consumption
	r := domain.NewReport()

	for id, ns := range nodesSimulations {
		for rf, as := range ns.RequestsFlows {
			for _, a := range as {
				for _, r := range a.Requirements {
					value, err := r.Calc(rf)
					if err != nil {
						logrus.WithFields(logrus.Fields{"nodeSimulation": ns, "requestsFlows": rf, "action": a}).Error(err)
						return nil, err
					}

					if nsr, ok := ns.Requirements[r.Resource]; ok {
						ns.Requirements[r.Resource] = r.Resource.Sum(nsr, value)
					} else {
						ns.Requirements[r.Resource] = value
					}
				}
			}
		}

		r.NodeReports[id] = domain.NewNodeReport(ns.Requirements)
	}

	return r, nil
}

func (suc *simulatorUsecase) requestsFlowGradientDescent(ns map[string]*domain.NodeSimulation, ls map[*domain.Link]struct{}, n *domain.Node, inrfs map[*domain.RequestsFlow][]*domain.Action) {
	// Add input requests flows
	ns[n.ID].RequestsFlows = inrfs

	// Create map with request flows for next nodes
	outrfs := make(map[*domain.RequestsFlow][]*domain.Action)

	// Translate input requests forward
	for rf := range inrfs {
		for _, link := range n.Links {
			if outrf, ok := outrfs[rf]; ok {
				outrf = append(outrf, link.Action)
			} else {
				outrfs[rf] = []*domain.Action{link.Action}
			}
		}
	}

	// Add this node requests to children
	if n.RequestsFlow != nil {
		for _, link := range n.Links {
			if outrf, ok := outrfs[n.RequestsFlow]; ok {
				outrf = append(outrf, link.Action)
			} else {
				outrfs[n.RequestsFlow] = []*domain.Action{link.Action}
			}
		}
	}

	// Go to children
	for _, link := range n.Links {
		// Check that link wasn't checked
		if _, ok := ls[link]; !ok {
			ls[link] = struct{}{}
			suc.requestsFlowGradientDescent(ns, ls, link.Child, outrfs)
		}
	}
}
