package domain

type Report struct {
	Roots []*NodeReport
}

type NodeReport struct {
	// CPU usage in MOPS
	CpuUsage float64
	// RAM usage in bytes
	RamUsage float64
	// Storage usage in bytes
	StorageUsage float64
	Links        []*LinkReport
}

type LinkReport struct {
	// Data flow in bytes
	DataFlow float64
}

func NewReport() *Report {
	return new(Report)
}
