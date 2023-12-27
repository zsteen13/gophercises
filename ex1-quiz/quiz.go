package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type QAPair struct {
	Question, Answer string
}

func main() {
	// Open the CSV file
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new reader
	reader := csv.NewReader(file)

	// Read the CSV data
	p, c := 0, 0
	for {
		record, err := reader.Read()
		if err != nil {
			// Handle end-of-file error
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading CSV:", err)
			return
		}

		// Handle invalid records
		if len(record) < 2 {
			fmt.Println("Invalid record:", record)
			continue
		}

		// Put record into QAPair object, increment problem number
		qa := QAPair{
			Question: record[0],
			Answer:   record[1],
		}
		p += 1

		// Print question and await user answer
		fmt.Printf("Problem #%d: %s = ", p, qa.Question)
		var input string
		fmt.Scanln(&input)

		// Increment correct count if user input matches answer
		if input == qa.Answer {
			c += 1
		}
	}

	// Print final score
	fmt.Printf("You scored %d out of %d.\n", c, p)
}
