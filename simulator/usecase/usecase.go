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
	r := domain.NewReport()

	// 1. Search all clients
	clients, err := suc.ccuc.CreateNodesChains(sc)
	if err != nil {
		return nil, err
	}

	for _, node := range clients {

	}

	// 1. Calc base consumption
	for _, node := range sc.Roots {
		switch node.Type {
		case domain.NodeType_Component:
			r.Roots = append(r.Roots, &domain.NodeReport{
				CpuUsage: node.Model,
			})

		case domain.NodeType_Custom:
			// TODO Implement
			logrus.Debug("skip consumption calc for Custom node")
		case domain.NodeType_Client:
			logrus.Debug("skip consumption calc for Client node")
		default:
			return nil, domain.ErrUnknownNodeType
		}
	}

	return r, nil
}
