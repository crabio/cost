package domain

type ResourceType string

const (
	ResourceType_Cpu                   ResourceType = "cpu"
	ResourceType_Ram                   ResourceType = "ram"
	ResourceType_StorageRead           ResourceType = "storage-read"
	ResourceType_StorageWrite          ResourceType = "storage-write"
	ResourceType_StorageReadPerByte    ResourceType = "storage-read-per-byte"
	ResourceType_StorageWritePerByte   ResourceType = "storage-write-per-byte"
	ResourceType_NetworkReceive        ResourceType = "network-receive"
	ResourceType_NetworkSend           ResourceType = "network-send"
	ResourceType_NetworkReceivePerByte ResourceType = "network-receive-per-byte"
	ResourceType_NetworkSendPerByte    ResourceType = "network-send-per-byte"
)
