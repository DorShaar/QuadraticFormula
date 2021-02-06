package main

import (
	"equationsolver"
	"rediscache"
	// "equationqueue"
	// "equationmessage"
	"log"
	"os"
	"github.com/google/uuid"
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
	redisCache := rediscache.RedisCache{}
	redisCache.Connect(redisConnectionAddress)

}