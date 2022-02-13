package domain

type RequestsFlow struct {
	Qty      uint `yaml:"qty"`
	PeriodMs uint `yaml:"period-ms"`
	MsgSize  uint `yaml:"msg-size"`
}
