package main

import (
	"equationsolver"
	"equationqueue"
	"equationmessage"
	"log"
	"os"
	"strconv"
)

const redisConnectionAddress = "localhost:6379"
const queueConnectionAddress = "localhost:61613"
const equationSolverQueueName = "equation-solver"
const equationReporterQueueName = "equation-reporter"

func main() {
	f, err := os.OpenFile("equationsolver/equation_solver.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	log.SetOutput(f)

	log.Println("Equation Solver is running")

	solveEquations()
}

func solveEquations() {
	equationSolverManager := equationsolver.EquationSolverManager{}
	equationSolverManager.Init(redisConnectionAddress)

	queueReader := equationqueue.QueueReader{}
	queueReader.Connect(queueConnectionAddress)
	queueReader.Subscribe(equationSolverQueueName)

	queueWriter := equationqueue.QueueWriter{}
	queueWriter.Connect(queueConnectionAddress)

	for true {
		equationMessage, err := queueReader.ReadMessage()

		if err != nil {
			log.Println("Error when trying to read message")
			continue
		}

		a, b, c, err := getCoefficientsFromMessage(equationMessage)

		if err != nil {
			log.Println("Could not get coefficients from message")
			continue
		}

		root1, root2, hasRoots := equationSolverManager.FindRoots(a, b, c)

		if hasRoots {
			equationMessage.Root1 = float32(root1)
			equationMessage.Root2 = float32(root2)
		} else {
			equationMessage.Root1 = 0
			equationMessage.Root2 = 0
		}

		queueWriter.SendMessage(equationMessage, equationReporterQueueName)
	}
}

func getCoefficientsFromMessage(equationMessage *equationmessage.EquationMessage) (int, int, int, error){
	a, err := strconv.Atoi(equationMessage.A)
	if err != nil {
		log.Println("Could not parse %s to int", equationMessage.A)
		return 0, 0, 0, err
	}

	b, err := strconv.Atoi(equationMessage.B)
	if err != nil {
		log.Println("Could not parse %s to int", equationMessage.B)
		return 0, 0, 0, err
	}

	c, err := strconv.Atoi(equationMessage.C)
	if err != nil {
		log.Println("Could not parse %s to int", equationMessage.C)
		return 0, 0, 0, err
	}

	return a, b, c, nil
}