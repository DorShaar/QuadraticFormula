package EquationReporter

import (
	"container/list"
	"time"
)

type Report struct {
	startTime           time.Time
	endingTime          time.Time
	equationMessageList list.List
}

func (report Report) WriteAllToCSV(csvPath string) {
	for message := report.equationMessageList.Front(); message != nil; message = message.Next() {
		// do something with e.Value
	}
}
