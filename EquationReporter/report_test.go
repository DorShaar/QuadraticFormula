package equationreporter

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	message := EquationMessage{
		correlationID:     "ABCD",
		originalEquation:  "x^2 + x^2 + x - 1 = 0",
		organizedEquation: "2x^2 + x - 1 = 0",
		a:                 5,
		b:                 6,
		c:                 6,
		root1:             1.221,
		root2:             3.45,
	}

	messages := []EquationMessage{
		message,
	}

	report := Report{
		endingTime:       time.Now(),
		startTime:        time.Now(),
		equationMessages: messages,
	}

	tempCSV := "temp.csv"

	if _, err := os.Stat(tempCSV); err == nil {
		t.FailNow()
	}

	report.WriteAllToCSV(tempCSV)
	defer deleteFile(tempCSV)

	if _, err := os.Stat(tempCSV); err != nil {
		t.FailNow()
	}
}

func deleteFile(fileName string) {
	e := os.Remove(fileName)

	if e != nil {
		log.Fatal(e)
	}
}
