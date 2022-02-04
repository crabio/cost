package usecase

import (
	"github.com/iakrevetkho/cost/domain"
	"github.com/sirupsen/logrus"
)

type SimulatorUsecase interface {
	Simulate(sc domain.SchemeConfig) (*domain.Report, error)
}

type simulatorUsecase struct {
}

func NewSimulatorUsecase() SimulatorUsecase {
	tuc := new(simulatorUsecase)
	return tuc
}

func (suc *simulatorUsecase) Simulate(sc domain.SchemeConfig) (*domain.Report, error) {
	r := domain.NewReport()

	// 1. Calc base consumption
	for _, node := range sc.Nodes {
		switch node.Type {
		case domain.NodeType_Component:
			node.Model

			sc.Models[]

		case domain.NodeType_Custom:
			// TODO Implement
			logrus.Debug("skip consumption calc for Custom node")
		case domain.NodeType_Client:
			logrus.Debug("skip consumption calc for Client node")
		default:
			return nil, domain.UNKNOWN_NODE_TYPE
		}
	}

	return r, nil
}
