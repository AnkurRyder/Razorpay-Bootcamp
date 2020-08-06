package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "../problem.csv"
	csvfile, err := os.Open(filename)

	if err != nil {
		fmt.Println("Could not open the problem file", err)
	}

	reader := csv.NewReader(csvfile)

	for {
		problem, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		println(problem[0])
	}

}
