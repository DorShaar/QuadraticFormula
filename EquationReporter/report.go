package equationreporter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

// Report hold the start and end time of a progress and a list of equation messages.
type Report struct {
	startTime        time.Time
	endingTime       time.Time
	equationMessages []EquationMessage
}

// WriteAllToCSV Writes report into given csvPath.
func (report Report) WriteAllToCSV(csvPath string, headers []string) {
	records := [][]string{
		headers,
	}

	for _, message := range report.equationMessages {
		rows := make([]string, len(records[0]))

		setCsvRows(rows, message)

		records = append(records, rows)

		writeToCsv(csvPath, records)
	}
}

func setCsvRows(rows []string, message EquationMessage) {
	i := 0
	rows[i] = message.correlationID
	i++
	rows[i] = message.originalEquation
	i++
	rows[i] = message.organizedEquation
	i++
	rows[i] = strconv.FormatInt(int64(message.a), 10)
	i++
	rows[i] = strconv.FormatInt(int64(message.b), 10)
	i++
	rows[i] = strconv.FormatInt(int64(message.c), 10)
	i++
	rows[i] = strconv.FormatFloat(float64(message.root1), 'f', -1, 64)
	i++

	if message.root1 != message.root2 {
		rows[i] = strconv.FormatFloat(float64(message.root2), 'f', -1, 64)
	}

	i++

	for _, tracedMethod := range message.tracedMethodsList {
		rows[i] = strconv.FormatInt(int64(tracedMethod.duration), 10)
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
