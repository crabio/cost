package usecase

import (
	"github.com/iakrevetkho/cost/domain"
)

type SimulatorUsecase interface {
	Simulate(dfc domain.DataFlowChain) (*domain.Report, error)
}

type simulatorUsecase struct {
}

func NewSimulatorUsecase() SimulatorUsecase {
	tuc := new(simulatorUsecase)
	return tuc
}

func (suc *simulatorUsecase) Simulate(dfc domain.DataFlowChain) (*domain.Report, error) {
	r := domain.NewReport()

	// // 1. Calc base consumption
	// for _, node := range sc.Roots {
	// 	switch node.Type {
	// 	case domain.NodeType_Component:
	// 		r.Roots = append(r.Roots, &domain.NodeReport{
	// 			CpuUsage: node.Model,
	// 		})

	// 	case domain.NodeType_Custom:
	// 		// TODO Implement
	// 		logrus.Debug("skip consumption calc for Custom node")
	// 	case domain.NodeType_Client:
	// 		logrus.Debug("skip consumption calc for Client node")
	// 	default:
	// 		return nil, domain.ErrUnknownNodeType
	// 	}
	// }

	return r, nil
}
