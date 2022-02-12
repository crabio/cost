package usecase

import (
	"github.com/iakrevetkho/cost/domain"
	"github.com/sirupsen/logrus"
)

type ChainCreatorUsecase interface {
	CreateChains(sc *domain.SchemeConfig) ([]*domain.DataFlowChain, error)
}

type chainCreatorUsecase struct {
}

func NewChainCreatorUsecase() ChainCreatorUsecase {
	ccuc := new(chainCreatorUsecase)
	return ccuc
}

func (ccuc *chainCreatorUsecase) CreateChains(sc *domain.SchemeConfig) ([]*domain.DataFlowChain, error) {
	// Create map with all nodes
	// Key - ID
	// Value - Node
	nodesMap := make(map[string]*domain.Node)
	// Create map with root nodes
	// Key - ID
	// Value - struct
	rootNodesMap := make(map[string]struct{})
	for id, node := range sc.Nodes {
		model, ok := sc.Models[node.Model]
		if !ok {
			logrus.WithField("node", node).Warn(domain.ErrUnknownModel)
		}

		logrus.WithField("node", node).Debug("node")

		nodeModel := domain.NewModel(node.Model, model.Name, model.Type, model.Params, model.AvailableActions)

		nodesMap[id] = domain.NewNode(id, node.Name, node.Type, nodeModel, []*domain.Link{})

		rootNodesMap[id] = struct{}{}
	}

	for linkId, link := range sc.Links {
		logrus.WithFields(logrus.Fields{"id": linkId, "link": link}).Debug("link")

		startNode, ok := nodesMap[link.Start]
		if !ok {
			logrus.WithField("link", link).Error(domain.ErrUnknownNodeId)
			return nil, domain.ErrUnknownNodeId
		}

		endNode, ok := nodesMap[link.End]
		if !ok {
			logrus.WithField("link", link).Error(domain.ErrUnknownNodeId)
			return nil, domain.ErrUnknownNodeId
		}

		action, ok := endNode.Model.AvailableActions[link.ActionName]
		if !ok {
			logrus.WithFields(logrus.Fields{"endNode": endNode, "link": link, "model": endNode.Model}).Error(domain.ErrUnknownModelAction)
			return nil, domain.ErrUnknownModelAction
		}

		switch action.Direction {
		case domain.Action_DirectionType_In:
			startNode.Links = append(startNode.Links, domain.NewLink(linkId, link.Seq, endNode, link.Type, &action))
			// Node is root if it has at least 1 out link (in for another component)
			delete(rootNodesMap, endNode.ID)
			logrus.WithFields(logrus.Fields{"action": action, "node": endNode}).Debug("delete root node")

		case domain.Action_DirectionType_Out:
			endNode.Links = append(endNode.Links, domain.NewLink(linkId, link.Seq, startNode, link.Type, &action))
			// Node is root if it has at least 1 out link (in for another component)
			delete(rootNodesMap, startNode.ID)
			logrus.WithFields(logrus.Fields{"action": action, "node": startNode}).Debug("delete root node")

		default:
			logrus.WithField("direction", action.Direction).Error(domain.ErrUnknownActionDirection)
			return nil, domain.ErrUnknownActionDirection
		}
	}

	var chains []*domain.DataFlowChain

	for id, _ := range rootNodesMap {
		logrus.WithField("id", id).Debug("root node")

		chain := domain.NewDataFlowChain(nil, nodesMap[id], nil, nil)

		chains = append(chains, chain)
	}

	return chains, nil
}
