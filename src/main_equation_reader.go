package main

import (
	"equationreader"
	"equationqueue"
	"log"
	"os"
)

const connectionAddress = "localhost:61613"
const publishQueueName = "equation-arranger"

func main() {
	f, err := os.OpenFile("EquationReader/equation_reader.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	for _, equation := range equations {
		queueWriter.SendMessage(publishQueueName, equation)
	}

	queueWriter.Disconnect()
}
