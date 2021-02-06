package main

import (
	"log"
	"os"
	"equationarranger"
	"equationscanner"
	"equationqueue"
	"equationmessage"
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
		equationMessage, err := queueReader.ReadMessage()

		if err != nil {
			log.Println("Error when trying to read message")
			continue
		}

		arrangeResult := equationArranger.Arrange(equationMessage.OriginalEquation)

		if arrangeResult.IsArrangeSucceeded() {
			equationMessage = &equationmessage.EquationMessage { 
				CorrelationId: equationMessage.CorrelationId,
				OriginalEquation: equationMessage.OriginalEquation,
				ArrangedEquation: arrangeResult.ArrangedEquation(),
			}

			queueWriter.SendMessage(equationMessage, equationDisassemblerQueueName)
		}
	}
}