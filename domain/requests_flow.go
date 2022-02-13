package domain

type RequestsFlow struct {
	Qty      uint `yaml:"qty"`
	PeriodMs uint `yaml:"period-ms"`
	MsgSize  uint `yaml:"msg-size"`
}

func (rf *RequestsFlow) RequestsPerSecond() uint {
	return rf.Qty * (1000 / rf.PeriodMs)
}
