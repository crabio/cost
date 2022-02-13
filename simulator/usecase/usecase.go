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
	for _, n := range clients {
		suc.requestsFlowGradientDescent(nodesSimulations, nil, n, nil, nil)
	}

	// 4. Go through all nodes and calc consumption
	r := domain.NewReport()

	for id, ns := range nodesSimulations {
		for rf, as := range ns.ActionsFlows {
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

func (suc *simulatorUsecase) requestsFlowGradientDescent(ns map[string]*domain.NodeSimulation, ls map[*domain.Link]struct{}, n *domain.Node, inrfs []*domain.RequestsFlow, ina *domain.Action) {
	if ls == nil {
		ls = make(map[*domain.Link]struct{})
	}

	// Add input requests flow
	ns[n.ID].RequestsFlows = inrfs
	// Add nodes requests flow to overall if has
	if n.RequestsFlow != nil {
		ns[n.ID].RequestsFlows = append(ns[n.ID].RequestsFlows, n.RequestsFlow)
	}

	// Add actions for current node multiplied on inpur requests flow
	for _, rf := range inrfs {
		if af, ok := ns[n.ID].ActionsFlows[rf]; ok {
			af = append(af, ina)
		} else {
			ns[n.ID].ActionsFlows[rf] = []*domain.Action{ina}
		}
	}

	// Go to children
	for _, link := range n.Links {
		// Check that link wasn't checked
		if _, ok := ls[link]; !ok {
			ls[link] = struct{}{}
			logrus.WithFields(logrus.Fields{"node": n, "child": link.Child}).Debug("go to child")
			suc.requestsFlowGradientDescent(ns, ls, link.Child, ns[n.ID].RequestsFlows, link.Action)
		}
	}
}
