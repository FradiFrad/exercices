package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var quizChan chan int
	var goodAnswers int

	// Parse the csv
	questions, answers := CSVParser("./problems.csv")

	totalQuestions := len(questions)

	timerDuration := addATimerFlag()

	// Reminder: waitgroup is mandatory to wait for a function to finished before launching a goroutine
	var wg sync.WaitGroup
	wg.Add(1)

	// WARNING Note that a WaitGroup must be passed to functions by pointer.
	startTimer(&wg, timerDuration)
	wg.Wait()

	// Launch the quiz in a gouroutine that puts good answers into a chan
	go handleAnswers(questions, answers, quizChan, totalQuestions)

	// Have to use timer instead of time to allow user to start it via input
	timer := time.NewTimer(time.Duration(*timerDuration * int(time.Second)))

	// is select a goroutine?
	// 2 chans :
	//	- one with the good answers : when the chan has it, it means the quiz is finished
	//	- one with the timer : when the chan is filled, it means the time is up
	select {
	case goodAnswers = <-quizChan:
	case <-timer.C:
		fmt.Print("Time's up ! \n")
	}
	// Give correct answers number and total questions number
	if goodAnswers > 5 {
		fmt.Printf("Congratulations ! You have %d good answers on %d questions !\n", goodAnswers, totalQuestions)
	}
	fmt.Printf("Oups ! You only have %d good answers on %d questions... Let's try again \n", goodAnswers, totalQuestions)
	return
}

// Allow the user to custom the timer via a flag (-timer) in CLI
// 	- Parse() is mandatory to take user input into account
// Cf. https://golang.org/pkg/flag/#pkg-examples
func addATimerFlag() *int {
	timerDuration := flag.Int("timer", 30, "duration in seconds")
	flag.Parse()
	return timerDuration
}

// Give instructions about the timer to the user
// 	- Scan() is mandatory to take user input into account
// Cf. https://golang.org/pkg/bufio
func startTimer(wg *sync.WaitGroup, timerDuration *int) {
	fmt.Print("Hi! You're about to start a very interesting quiz...\n But you'll only have 30s to answer all the questions ! Ready?\n [Press any key to begin] or \n [Change the timer by launching the quiz again with 'go run main.go -timer=yourDuration' (in seconds)]\n")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Print("You didn't answer  \n")
	} else {
		wg.Done()
	}
}

// Launch the quiz questions and put goodAnswers into a chan
func handleAnswers(questions map[int]string, answers map[int]string, quizChan chan<- int, totalQuestions int) {
	goodAnswers := 0
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
}

// CSVParser parse the source file for the quiz
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
