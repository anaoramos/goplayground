package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type CsvDataLines struct {
	Column1 string
	Column2 string
}

func ReadCsvFile(filename string) ([][]string, error) {
	// Open CSV file
	fileContent, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer fileContent.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func StarQuiz(csvData [][]string) {
	correctAnswers := 0
	incorrectAnswers := 0

	reader := bufio.NewReader(os.Stdin)
	for _, line := range csvData {
		data := CsvDataLines{
			Column1: line[0],
			Column2: line[1],
		}

		fmt.Printf("Q: %s (y/n)", data.Column1) // Question
		text, _ := reader.ReadString('\n')
		if data.Column2 == strings.TrimSpace(text) {
			correctAnswers += 1
			fmt.Println("Correct! :) \n")
		} else {
			incorrectAnswers += 1
			fmt.Println("Incorrect! :(")
		}
		fmt.Printf("Correct Answers: %d\nIncorrect Answers: %d\n", correctAnswers, incorrectAnswers)
	}

	if correctAnswers > incorrectAnswers {
		fmt.Println("Winner!")
	} else {
		fmt.Println("Loser!")
	}
}

func main() {
	// Read the CSV file from command-line argument
	if len(os.Args) < 2 {
		log.Fatal("CSV argument is required")
	}
	csvFile := os.Args[1]
	csvData, err := ReadCsvFile(csvFile)
	if err != nil {
		panic(err)
	}
	StarQuiz(csvData)

}
