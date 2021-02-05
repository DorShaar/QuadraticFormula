package main

import (
	"log"
	"os"
	"equationarranger"
	"equationscanner"
	"equationqueue"
)

const connectionAddress = "localhost:61613"
const equationArrangerQueueName = "equation-arranger"
const equationDisassemblerQueueName = "equation-disassembler"

func main() {
	f, err := os.OpenFile("equationarranger/equation_arranger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	log.SetOutput(f)

	log.Println("Equation Arranger is running")

	arrangeEquations()
}

func arrangeEquations() {
	queueReader := equationqueue.QueueReader{}
	queueReader.Connect(connectionAddress)
	queueReader.Subscribe(equationArrangerQueueName)

	queueWriter := equationqueue.QueueWriter{}
	queueWriter.Connect(connectionAddress)

	equationScanner := equationscanner.EquationScanner{}
	equationArranger := equationarranger.EquationArranger{EquationScanner: equationScanner}

	for true {
		equation := queueReader.ReadMessage()
		arrangeResult := equationArranger.Arrange(equation)

		if arrangeResult.IsArrangeSucceeded() {
			queueWriter.SendMessage(equationDisassemblerQueueName, arrangeResult.ArrangedEquation())
		}
	}
}