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
func (report Report) WriteAllToCSV(csvPath string) {
	records := [][]string{
		{"correlation ID", "original equation", "organized equation", "a", "b", "c", "root1", "root2", "methods"},
	}

	for _, message := range report.equationMessages {
		row := make([]string, 9)

		row[0] = message.correlationID
		row[1] = message.originalEquation
		row[2] = message.organizedEquation
		row[3] = strconv.FormatInt(int64(message.a), 10)
		row[4] = strconv.FormatInt(int64(message.b), 10)
		row[5] = strconv.FormatInt(int64(message.c), 10)
		row[6] = strconv.FormatFloat(float64(message.root1), 'f', -1, 64)

		if message.root1 != message.root2 {
			row[7] = strconv.FormatFloat(float64(message.root2), 'f', -1, 64)
		}

		row[8] = " "

		records = append(records, row)

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
}
