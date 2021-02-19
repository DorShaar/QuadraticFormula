package equationreportertests

import (
	"log"
	"os"
	"testing"
	"equationreporter"
	"equationmessage"
)

func TestSum(t *testing.T) {
	tracedMethod1 := equationmessage.TracedMethod{
		MethodName: "split numbers",
		Duration:   100,
	}

	tracedMethod2 := equationmessage.TracedMethod{
		MethodName: "find parameters",
		Duration:   161,
	}

	tracedMethod3 := equationmessage.TracedMethod{
		MethodName: "long operation",
		Duration:   1306,
	}

	tracedMethods := make([]equationmessage.TracedMethod, 0)

	tracedMethods = append(tracedMethods, tracedMethod1)
	tracedMethods = append(tracedMethods, tracedMethod2)
	tracedMethods = append(tracedMethods, tracedMethod3)

	message := equationmessage.EquationMessage{
		CorrelationId:     "ABCD",
		OriginalEquation:  " x^2 + x^2 + x - 1 =0",
		ArrangedEquation: "2x^2 + x - 1 = 0",
		A:                 "5",
		B:                 "6",
		C:                 "6",
		Root1:             1.221,
		Root2:             3.45,
		TracedMethodsList: tracedMethods,
	}

	report := equationreporter.EquationReporter{}

	report.AddMessage(&message)

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

func createHeaders(tracedMethods []equationmessage.TracedMethod) []string {
	headers := make([]string, 0)

	headers = append(headers, "Correlation Id")
	headers = append(headers, "Original equation")
	headers = append(headers, "Arranged equation")
	headers = append(headers, "A")
	headers = append(headers, "B")
	headers = append(headers, "C")
	headers = append(headers, "Root1")
	headers = append(headers, "Root2")

	for _, tracedMethod := range tracedMethods {
		headers = append(headers, tracedMethod.MethodName)
	}

	return headers
}

func deleteFile(fileName string) {
	e := os.Remove(fileName)

	if e != nil {
		log.Fatal(e)
	}
}