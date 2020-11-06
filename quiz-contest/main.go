package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// opening csv file
	file := "problems.csv"
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// reading csv file
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// variable to keep a total count of correct answers
	correctCount := 0

	for i := 0; i < len(records); i++ {
		// Printing the question
		fmt.Println("What is the answer of: ", records[i][0])

		// taking answer input as int
		var answer int
		fmt.Scan(&answer)

		// convert string read from csv to int to compare later with answer
		actualAnswer, err := strconv.Atoi(records[i][1])
		if err != nil {
			log.Fatal(err)
		}

		// increment correct counter if the answer is correct.
		if answer == actualAnswer {
			correctCount = correctCount + 1
		}
	}
	fmt.Println("total correct answers are: ", correctCount)
}
