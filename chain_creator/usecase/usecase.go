package usecase

// TODO Chain Creator won;t be used now. It will be used for creating simulation with lags calc

import (
	"github.com/iakrevetkho/cost/domain"
	"github.com/sirupsen/logrus"
)

type ChainCreatorUsecase interface {
	CreateNodesChains(sc *domain.SchemeConfig) ([]*domain.Node, error)
	CreateChains(sc *domain.SchemeConfig) ([]*domain.DataFlowChainBlock, error)
}

type chainCreatorUsecase struct {
}

func NewChainCreatorUsecase() ChainCreatorUsecase {
	ccuc := new(chainCreatorUsecase)
	return ccuc
}

func (ccuc *chainCreatorUsecase) CreateNodesChains(sc *domain.SchemeConfig) ([]*domain.Node, error) {
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

		if node.Type == domain.NodeType_Client {
			nodesMap[id] = domain.NewClientNode(id, node.Name, node.Type, nodeModel, []*domain.Link{}, node.RequestsFlow)
		} else {
			nodesMap[id] = domain.NewNode(id, node.Name, node.Type, nodeModel, []*domain.Link{})
		}

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

	var rootNodes []*domain.Node

	for id := range rootNodesMap {
		logrus.WithField("id", id).Debug("root node")

		rootNodes = append(rootNodes, nodesMap[id])
	}

	return rootNodes, nil
}

func (ccuc *chainCreatorUsecase) CreateChains(sc *domain.SchemeConfig) ([]*domain.DataFlowChainBlock, error) {
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

	var chains []*domain.DataFlowChainBlock

	for id := range rootNodesMap {
		logrus.WithField("id", id).Debug("root node")

		chain := domain.NewDataFlowChainBlock(nil, nodesMap[id], nil, nil)

		// node, ok := nodesMap[id]
		// if !ok {
		// 	return nil, domain.ErrUnknownNodeId
		// }

		// // Sort links by seq
		// sort.Slice(node.Links, func(i, j int) bool {
		// 	return node.Links[i].Seq < node.Links[j].Seq
		// })

		// for _, link := range node.Links {

		// }

		chains = append(chains, chain)
	}

	return chains, nil
}

func (ccuc *chainCreatorUsecase) CreateChain(node *domain.Node, prev *domain.DataFlowChainBlock) (*domain.DataFlowChainBlock, error) {
	// Seeking for link that not in previous chain's blocks
	for _, link := range node.Links {
		if prev != nil {
			if prev.IsLinkExists(link) {
				// Skip chained links
				continue
			}
		}

		// 1. Go forward by all links with action "in" type
		if link.Action.Direction == domain.Action_DirectionType_In {
			block := domain.NewDataFlowChainBlock(prev, node, link.Child, link)

			return ccuc.CreateChain(node, block)
		} else {
			// TODO
		}
	}
	return nil, nil
}
