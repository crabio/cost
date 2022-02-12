package domain

type Requirement struct {
	Resource      RequirementResourceType `yaml:"resource"`
	Type          RequirementType         `yaml:"type"`
	UnitOfMeasure UnitOfMeasure           `yaml:"uom"`
	Value         float64                 `yaml:"value"`
}

type RequirementResourceType string

const (
	RequirementResourceType_Cpu                   RequirementResourceType = "cpu"
	RequirementResourceType_Ram                   RequirementResourceType = "ram"
	RequirementResourceType_StorageRead           RequirementResourceType = "storage-read"
	RequirementResourceType_StorageWrite          RequirementResourceType = "storage-write"
	RequirementResourceType_StorageReadPerByte    RequirementResourceType = "storage-read-per-byte"
	RequirementResourceType_StorageWritePerByte   RequirementResourceType = "storage-write-per-byte"
	RequirementResourceType_NetworkReceive        RequirementResourceType = "network-receive"
	RequirementResourceType_NetworkSend           RequirementResourceType = "network-send"
	RequirementResourceType_NetworkReceivePerByte RequirementResourceType = "network-receive-per-byte"
	RequirementResourceType_NetworkSendPerByte    RequirementResourceType = "network-send-per-byte"
)

type RequirementType string

const (
	RequirementType_Once           RequirementType = "once"
	RequirementType_PerRequest     RequirementType = "per-request"
	RequirementType_PerRequestByte RequirementType = "per-request-byte"
)
