package domain

import (
	"math"

	"github.com/sirupsen/logrus"
)

type ResourceType string

const (
	ResourceType_Cpu            ResourceType = "cpu"
	ResourceType_Ram            ResourceType = "ram"
	ResourceType_StorageRead    ResourceType = "storage-read"
	ResourceType_StorageWrite   ResourceType = "storage-write"
	ResourceType_NetworkReceive ResourceType = "network-receive"
	ResourceType_NetworkSend    ResourceType = "network-send"
)

func (rt ResourceType) Uom() UnitOfMeasure {
	switch rt {
	case ResourceType_Cpu:
		return UnitOfMeasure_Ops

	case ResourceType_Ram,
		ResourceType_StorageRead,
		ResourceType_StorageWrite,
		ResourceType_NetworkReceive,
		ResourceType_NetworkSend:
		return UnitOfMeasure_Byte

	default:
		logrus.WithField("type", rt).Fatal(ErrUnknownResourceType)
		return UnitOfMeasure_NA
	}
}

func (rt ResourceType) Sum(a, b float64) float64 {
	switch rt {
	case ResourceType_Cpu,
		ResourceType_StorageRead,
		ResourceType_StorageWrite,
		ResourceType_NetworkReceive,
		ResourceType_NetworkSend:
		return a + b

	case ResourceType_Ram:
		return math.Max(a, b)

	default:
		logrus.WithField("type", rt).Fatal(ErrUnknownResourceType)
		return 0
	}
}
