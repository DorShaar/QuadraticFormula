package equationreporter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
	"equationmessage"
)

// Report holds the start and end time of a progress and a list of equation messages.
type EquationReporter struct {
	startTime        time.Time
	endingTime       time.Time
	equationMessages []equationmessage.EquationMessage
}

// WriteAllToCSV Writes report into given csvPath.
func (report *EquationReporter) WriteAllToCSV(csvPath string, headers []string) {
	records := [][]string{
		headers,
	}

	report.endingTime = time.Now()

	for _, message := range report.equationMessages {
		rows := make([]string, len(records[0]))

		setCsvRows(rows, &message)

		records = append(records, rows)

		writeToCsv(csvPath, records)
	}
}

func setCsvRows(rows []string, message *equationmessage.EquationMessage) {
	i := 0
	rows[i] = message.CorrelationId
	i++
	rows[i] = message.OriginalEquation
	i++
	rows[i] = message.ArrangedEquation
	i++
	rows[i] = message.A
	i++
	rows[i] = message.B
	i++
	rows[i] = message.C
	i++
	rows[i] = strconv.FormatFloat(float64(message.Root1), 'f', -1, 32)
	i++

	if message.Root1 != message.Root2 {
		rows[i] = strconv.FormatFloat(float64(message.Root2), 'f', -1, 32)
	}

	i++

	for _, tracedMethod := range message.TracedMethodsList {
		rows[i] = strconv.FormatInt(int64(tracedMethod.Duration), 10)
		i++
	}
}

func writeToCsv(csvPath string, records [][]string) {
	csvFile, err := os.Create(csvPath)
	defer csvFile.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	for _, record := range records {
		err := csvWriter.Write(record)
		if err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}

func (report *EquationReporter) AddMessage(message *equationmessage.EquationMessage) {
	if len(report.equationMessages) == 0 {
		report.startTime = time.Now()
	}

	report.equationMessages = append(report.equationMessages, *message)
}