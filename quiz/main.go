package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"io/ioutil"
	"log"
	"strings"
)

func main()  {
	goodAnswers := 0
	
	// Parse the csv
	questions, answers := CSVParser("./problems.csv")

	// Ask the questions, collect good answers
	totalQuestions := len(questions)
	for i := 1; i <= len(questions); i++ {
		fmt.Printf("Question %d : %s = ? \n", i, questions[i])
		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if answer == answers[i] {
			goodAnswers++
		}
	}

	// Give correct answers number and total questions number
	if goodAnswers > 5 {
		fmt.Printf("Congratulations ! You have %d good answers on %d questions !\n", goodAnswers, totalQuestions)
		return
	}
	fmt.Printf("Oups ! You only have %d good answers on %d questions... Let's try again \n", goodAnswers, totalQuestions)
	
	return
}

func CSVParser(path string) (map[int]string, map[int]string)  {
	questions := make(map[int]string)
	answers := make(map[int]string)
	i := 1

	// Take the CSV content
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	
	// Parse the CSV content
	CSVreader := csv.NewReader(strings.NewReader(string(content)))

	for {
		record, err := CSVreader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		questions[i] = record[0]
		answers[i] = record[1]

		i++
	}
	
	return questions, answers
}