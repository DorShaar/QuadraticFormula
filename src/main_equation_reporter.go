package main

import (
	"equationreporter"
	"equationqueue"
	"equationmessage"
	"log"
	"os"
	"time"
	"strings"
	"path/filepath"
)

const reportsDirectoryPath = "reports"
const queueConnectionAddress = "localhost:61613"
const equationReporterQueueName = "equation-reporter"

func main() {
	f, err := os.OpenFile("equationreporter/equation_reporter.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	log.SetOutput(f)

	log.Println("Equation Reporter is running")

	reportEquations()
}

func reportEquations() {
	equationReporter := equationreporter.EquationReporter{}

	queueReader := equationqueue.QueueReader{}
	queueReader.Connect(queueConnectionAddress)
	queueReader.Subscribe(equationReporterQueueName)

	var lastCorrelationIdSignal string
	var lastCorrelationId string

	for true {
		equationMessage, err := queueReader.ReadMessage()

		if err != nil {
			log.Println("Error when trying to read message")
			continue
		}

		isStopSignal := checkIfStopSignal(equationMessage)

		if isStopSignal  {
			lastCorrelationIdSignal = equationMessage.ArrangedEquation
			log.Printf("Signaled with last correlation id %s", lastCorrelationIdSignal)
		} else {
			lastCorrelationId = equationMessage.CorrelationId
			equationReporter.AddMessage(equationMessage)	
		}

		if lastCorrelationIdSignal != "" && lastCorrelationIdSignal == lastCorrelationId {
			break
		}
	}

	equationReporter.WriteAllToCSV(getReportPath(), createHeaders())
}

func checkIfStopSignal(equationMessage *equationmessage.EquationMessage) bool {
	return equationMessage.OriginalEquation == "end"
}

func createHeaders() []string {
	tracedMethods := make([]equationmessage.TracedMethod, 0)

	headers := make([]string, 0)

	headers = append(headers, "Correlation Id")
	headers = append(headers, "Original equation")
	headers = append(headers, "Arranged equation")
	headers = append(headers, "A")
	headers = append(headers, "B")
	headers = append(headers, "C")
	headers = append(headers, "Root1")
	headers = append(headers, "Root2")

	for _, tracedMethod := range tracedMethods {
		headers = append(headers, tracedMethod.MethodName)
	}

	return headers
}

func getReportPath() string {
	t := time.Now()
	reportName := t.Format("01-02-2006 15:04:05")
	validReportName := strings.ReplaceAll(reportName, ":", "_")
	reportPath := filepath.Join(reportsDirectoryPath, validReportName)

	return reportPath
}