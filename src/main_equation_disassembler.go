package main

import (
	"equationdisassembler"
	"equationqueue"
	"equationmessage"
	"log"
	"os"
)

const connectionAddress = "localhost:61613"
const equationDisassemblerQueueName = "equation-disassembler"
const equationSolverQueueName = "equation-solver"

func main() {
	f, err := os.OpenFile("equationdisassembler/equation_disassembler.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	log.SetOutput(f)

	log.Println("Equation Disassembler is running")

	disassembleEquations()
}

func disassembleEquations() {
	queueReader := equationqueue.QueueReader{}
	queueReader.Connect(connectionAddress)
	queueReader.Subscribe(equationDisassemblerQueueName)

	queueWriter := equationqueue.QueueWriter{}
	queueWriter.Connect(connectionAddress)

	for true {
		equationMessage, err := queueReader.ReadMessage()

		if err != nil {
			log.Println("Error when trying to read message")
			continue
		}

		disassembleResult := equationdisassembler.Disassemble(equationMessage.ArrangedEquation, "x")

		if !disassembleResult.IsDisassembleFailed {
			equationMessage = &equationmessage.EquationMessage { 
				CorrelationId: equationMessage.CorrelationId,
				OriginalEquation: equationMessage.OriginalEquation,
				ArrangedEquation: equationMessage.ArrangedEquation,
				A: disassembleResult.A,
				B: disassembleResult.B,
				C: disassembleResult.C,
			}

			queueWriter.SendMessage(equationMessage, equationSolverQueueName)
		}
	}
}