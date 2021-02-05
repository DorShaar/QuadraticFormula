package main

import (
	"equationdisassembler"
	"equationqueue"
	"log"
	"os"
)

const connectionAddress = "localhost:61613"
const equationDisassemblerQueueName = "equation-disassembler"
const equationSolverQueueName = "equation-solver"

func main() {
	f, err := os.OpenFile("equationarranger/equation_arranger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
		message := queueReader.ReadMessage()
		disassembleResult := equationdisassembler.Disassemble(message, message) // TODO

		if !disassembleResult.IsDisassembleFailed {
			queueWriter.SendMessage(equationSolverQueueName, message)
		}
	}
}