package equationreader

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

// ReadCsv reads from csv path.
func ReadCsv(path string) []string {
	csvfile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	equations := make([]string, 0)

	csvReader := csv.NewReader(csvfile)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			break
		}

		if !minimalEquationValidate(record[0]) {
			continue
		}

		equations = append(equations, record[0])
	}

	log.Printf("Found %d equations to solve", len(equations))
	return equations
}

func minimalEquationValidate(equationCandidate string) bool {
	return strings.ContainsAny(equationCandidate, "+-")
}
