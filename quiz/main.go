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

	// Parse the csv
	questions, answers := CSVParser("./problems.csv")
	goodAnswers := 0
	totalQuestions := len(questions)

	// Add a timer, by default 30s
	//timer := time.NewTimer(2 * time.Second)
	//	var timerChan chan time.Duration
	var quizChan chan string

	// Stop the timer at the end of the function.
	// Defers are called when the parent function exits.
	//defer timer.Stop()
	// t := time.Now()
	// fmt.Printf("Temps restant : %d", int(time.Until(t.Add(30*time.Second))))
	// go func() {
	// 	timerChan <- time.Until(t.Add(30 * time.Second))
	// }()
	for {
		go handleAnswers(questions, answers, quizChan)

		select {
		case <-quizChan:
			continue
		case <-time.After(30 * time.Second):
			fmt.Print("Time's up ! \n")
			return

		}

	}
	if goodAnswers > 5 {
		fmt.Printf("Congratulations ! You have %d good answers on %d questions !\n", goodAnswers, totalQuestions)
		return
	}
	fmt.Printf("Oups ! You only have %d good answers on %d questions... Let's try again \n", goodAnswers, totalQuestions)

}

// Ask the questions, count good answers
func handleAnswers(questions map[int]string, answers map[int]string, quizChan chan string) int {
	goodAnswers := 0

	totalQuestions := len(questions)
	// fmt.Print(remainingTime)

	for i := 1; i <= totalQuestions; i++ {
		fmt.Printf("Question %d : %s = ? \n", i, questions[i])
		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {
			fmt.Printf("You didn't answer question %d \n", i)
		}
		// Questions given invalid answers or unanswered are considered incorrect.
		if answer == answers[i] {
			goodAnswers++
		}
		quizChan <- answer
	}
	return goodAnswers
	// Give correct answers number and total questions number

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
