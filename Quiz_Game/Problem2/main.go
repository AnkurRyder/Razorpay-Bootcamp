package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type user struct {
	name                    string
	totalQuestionAttemted   int
	questionAnsweredCorrect int
}

func main() {
	filename := flag.String("filename", "../problem.csv", "a string")
	time := flag.String("time", "30s", "a string")
	flag.Parse()
	totalTime := *time
	reader := readproblem(*filename)
	var person user
	person.userinput()
	person.quiz(reader, totalTime)
	person.printResult()
}

func (person *user) userinput() {
	fmt.Println("please enter your name:")
	fmt.Scanln(&person.name)
	fmt.Println("please enter number of questions want to answer <= 12:")
	fmt.Scanln(&person.totalQuestionAttemted)
	person.questionAnsweredCorrect = 0
}

func readproblem(filename string) *csv.Reader {
	csvfile, err := os.Open(filename)

	if err != nil {
		fmt.Println("Could not open the problem file", err)
		os.Exit(1)
	}

	reader := csv.NewReader(csvfile)

	return reader
}

func (person *user) quiz(reader *csv.Reader, totalTime string) {
	i := 1
	var ans string
	c := make(chan string)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	go timer(c, totalTime)

	for {
		problem, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Q", i, ": ", problem[0])

		go takeInput(c)
		ans = <-c
		if ans == "EOT" {
			break
		}

		if ans == problem[1] {
			person.questionAnsweredCorrect = person.questionAnsweredCorrect + 1
		}

		if i == person.totalQuestionAttemted {
			break
		}
		i++
	}
}

func timer(c chan string, total string) {
	totalTime, _ := time.ParseDuration(total)
	start := time.Now()
	for {
		t := time.Now()
		elapsed := t.Sub(start)
		if elapsed >= totalTime {
			break
		}
	}
	fmt.Println("End of Time")
	c <- "EOT"
}

func takeInput(c chan string) {
	var ans string
	fmt.Scanln(&ans)
	c <- ans
}

func (person user) printResult() {
	fmt.Println(person.name, "attemted", person.totalQuestionAttemted, "Question and got",
		person.questionAnsweredCorrect, "answers correct")
}
