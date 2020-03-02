package main

import (
	"bufio"
	"encoding/csv"
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
	// Parse the csv
	questions, answers := CSVParser("./problems.csv")
	totalQuestions := len(questions)
	var quizChan chan int
	var goodAnswers int

	// Ping toutes les secondes
	// utiles pour shox les secondes restantes?
	// for range time.Tick(time.Second) {
	// 	println("ping!")
	// }
	// func elapsed(what string) func() {
	// 	start := time.Now()
	// 	return func() {
	// 		fmt.Printf("%s took %v\n", what, time.Since(start))
	// 	}
	// }

	// func main() {
	// 	defer elapsed("page")()  // <-- The trailing () is the deferred call
	// 	time.Sleep(time.Second * 2)
	// }
	// Rappel: waitgroup pour attendre qqch avant executio d'un routine
	var wg sync.WaitGroup
	wg.Add(1)

	// WARNING Note that a WaitGroup must be passed to functions by pointer.
	startTimer(&wg)
	wg.Wait()

	go handleAnswers(questions, answers, quizChan, totalQuestions)

	// Have to use timer instead of time to allow user to start it via input
	timer := time.NewTimer(3 * time.Second)
	// is select a goroutine?
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

func startTimer(wg *sync.WaitGroup) {
	fmt.Print("Hi! You're about to start a very interesting quiz... But you'll only have 30s to answer all the questions ! Ready? [Press any key to begin]\n")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Print("You didn't answer  \n")
	} else {
		wg.Done()
	}
}

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
