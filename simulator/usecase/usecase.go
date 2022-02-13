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
		suc.requestsFlowGradientDescent(nodesSimulations, nil, nil, n, nil, nil)
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

func (suc *simulatorUsecase) requestsFlowGradientDescent(ns map[string]*domain.NodeSimulation, ls map[*domain.Link]struct{}, parentNode, node *domain.Node, inrfs []*domain.RequestsFlow, ina *domain.Action) {
	if ls == nil {
		ls = make(map[*domain.Link]struct{})
	}

	// Add input requests flow
	ns[node.ID].RequestsFlows = inrfs
	// Add nodes requests flow to overall if has
	if node.RequestsFlow != nil {
		ns[node.ID].RequestsFlows = append(ns[node.ID].RequestsFlows, node.RequestsFlow)
	}

	// Add actions for current node multiplied on inpur requests flow
	if ina != nil {
		switch ina.Direction {
		case domain.Action_DirectionType_In:
			for _, rf := range inrfs {
				if af, ok := ns[node.ID].ActionsFlows[rf]; ok {
					af = append(af, ina)
				} else {
					ns[node.ID].ActionsFlows[rf] = []*domain.Action{ina}
				}
			}

		case domain.Action_DirectionType_Out:
			for _, rf := range inrfs {
				if af, ok := ns[parentNode.ID].ActionsFlows[rf]; ok {
					af = append(af, ina)
				} else {
					ns[parentNode.ID].ActionsFlows[rf] = []*domain.Action{ina}
				}
			}

		default:
			logrus.WithField("direction", ina.Direction).Error(domain.ErrUnknownActionDirection)
			return
		}
	}

	// Go to children
	for _, link := range node.Links {
		// Check that link wasn't checked
		if _, ok := ls[link]; !ok {
			ls[link] = struct{}{}
			suc.requestsFlowGradientDescent(ns, ls, node, link.Child, ns[node.ID].RequestsFlows, link.Action)
		}
	}
}
