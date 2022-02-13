package domain

import (
	"bytes"
	"strconv"
)

type RequestsFlow struct {
	Qty      uint `yaml:"qty"`
	PeriodMs uint `yaml:"period-ms"`
	MsgSize  uint `yaml:"msg-size"`
}

func (rf *RequestsFlow) RequestsPerSecond() uint {
	return rf.Qty * (1000 / rf.PeriodMs)
}

func (rf *RequestsFlow) String() string {
	var buf bytes.Buffer

	buf.WriteByte('{')
	buf.WriteString("qty: ")
	buf.WriteString(strconv.FormatUint(uint64(rf.Qty), 10))
	buf.WriteString(", period-ms: ")
	buf.WriteString(strconv.FormatUint(uint64(rf.PeriodMs), 10))
	buf.WriteString(", msg-size: ")
	buf.WriteString(strconv.FormatUint(uint64(rf.MsgSize), 10))
	buf.WriteByte('}')

	return buf.String()
}
