package main

import (
	"equationreader"
	"equationqueue"
	"equationmessage"
	"log"
	"os"
	"github.com/google/uuid"
)

const connectionAddress = "localhost:61613"
const equationArrangerQueueName = "equation-arranger"
const equationReporterQueueName = "equation-reporter"

func main() {
	f, err := os.OpenFile("equationreader/equation_reader.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	log.SetOutput(f)

	log.Println("Equation Reader is running")

	equations := readEquations()
	
	sendEquations(equations)
}

func readEquations() []string {
	equationsCsvPath := "../Resources/Equations.csv"
	log.Printf("Reading from %s", equationsCsvPath)
	equations := equationreader.ReadCsv(equationsCsvPath)
	return equations
}

func sendEquations(equations []string) {
	queueWriter := equationqueue.QueueWriter{}

	queueWriter.Connect(connectionAddress)

	var lastCorrelationId string
	for _, equation := range equations {
		equationMessage := &equationmessage.EquationMessage { 
				CorrelationId: uuid.New().String(),
				OriginalEquation: equation,
			}

		lastCorrelationId = equationMessage.CorrelationId

		queueWriter.SendMessage(equationMessage, equationArrangerQueueName)
	}

	equationMessage := &equationmessage.EquationMessage { 
			CorrelationId: uuid.New().String(),
			OriginalEquation: "end",
			ArrangedEquation: lastCorrelationId,
		}

	queueWriter.SendMessage(equationMessage, equationReporterQueueName)
	log.Printf("Sending end message - last correlation id is %s", lastCorrelationId)

	queueWriter.Disconnect()
}