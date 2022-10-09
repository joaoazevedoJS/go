package CSV

import (
	"encoding/csv"
	"log"
	"os"
)

type CreateCSV struct {
	Filename string
	Titles   []string
	Rows     [][]string
}

func Create(createCSV CreateCSV) {
	file, err := os.Create(createCSV.Filename + ".csv")

	defer file.Close()

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	Writer(csv.NewWriter(file), createCSV)
}

func Writer(writer *csv.Writer, createCSV CreateCSV) {
	defer writer.Flush()

	_ = writer.Write(createCSV.Titles)
	_ = writer.WriteAll(createCSV.Rows)

	println("Finish download")
}
