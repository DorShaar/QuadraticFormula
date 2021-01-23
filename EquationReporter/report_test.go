package equationreporter

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	tracedMethod1 := TracedMethod{
		methodName: "split numbers",
		duration:   100,
	}

	tracedMethod2 := TracedMethod{
		methodName: "find parameters",
		duration:   161,
	}

	tracedMethod3 := TracedMethod{
		methodName: "long operation",
		duration:   1306,
	}

	tracedMethods := make([]TracedMethod, 0)

	tracedMethods = append(tracedMethods, tracedMethod1)
	tracedMethods = append(tracedMethods, tracedMethod2)
	tracedMethods = append(tracedMethods, tracedMethod3)

	message := EquationMessage{
		correlationID:     "ABCD",
		originalEquation:  " x^2 + x^2 + x - 1 =0",
		organizedEquation: "2x^2 + x - 1 = 0",
		a:                 5,
		b:                 6,
		c:                 6,
		root1:             1.221,
		root2:             3.45,
		tracedMethodsList: tracedMethods,
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

	headers := createHeaders(tracedMethods)

	report.WriteAllToCSV(tempCSV, headers)
	defer deleteFile(tempCSV)

	if _, err := os.Stat(tempCSV); err != nil {
		t.FailNow()
	}
}

func createHeaders(tracedMethods []TracedMethod) []string {
	headers := make([]string, 0)

	headers = append(headers, "correlation ID")
	headers = append(headers, "original equation")
	headers = append(headers, "organized equation")
	headers = append(headers, "a")
	headers = append(headers, "b")
	headers = append(headers, "c")
	headers = append(headers, "root1")
	headers = append(headers, "root2")

	for _, tracedMethod := range tracedMethods {
		headers = append(headers, tracedMethod.methodName)
	}

	return headers
}

func deleteFile(fileName string) {
	e := os.Remove(fileName)

	if e != nil {
		log.Fatal(e)
	}
}
