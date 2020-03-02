package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	var quizChan chan int
	goodAnswers := 0

	// Parse the csv
	questions, answers := CSVParser("./problems.csv")

	// Ask the questions, collect good answers
	totalQuestions := len(questions)

	// Ping toutes les secondes
	// utiles pour shox les secondes restantes?
	// for range time.Tick(time.Second) {
	// 	println("ping!")
	// }
	go func() {
		for i := 1; i <= totalQuestions; i++ {
			fmt.Printf("Question %d : %s = ? \n", i, questions[i])
			var answer string
			_, err := fmt.Scanln(&answer)
			if err != nil {
				fmt.Printf("You didn't answer question %d \n", i)
			}
			if answer == answers[i] {
				goodAnswers++
			}
		}
		quizChan <- goodAnswers
	}()

	// is select a goroutine?
	select {
	case goodAnswers = <-quizChan:
		// Give correct answers number and total questions number
		if goodAnswers > 5 {
			fmt.Printf("Congratulations ! You have %d good answers on %d questions !\n", goodAnswers, totalQuestions)
		}
		fmt.Printf("Oups ! You only have %d good answers on %d questions... Let's try again \n", goodAnswers, totalQuestions)
		return

	case <-time.After(2 * time.Second):
		fmt.Print("Time's up ! \n")
		return
	}
}

func CSVParser(path string) (map[int]string, map[int]string) {
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
