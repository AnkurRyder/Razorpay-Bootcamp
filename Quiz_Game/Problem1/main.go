package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

type user struct {
	name                    string
	totalQuestionAttemted   int
	questionAnsweredCorrect int
}

func main() {
	filename := flag.String("filename", "../problem.csv", "a string")
	flag.Parse()
	reader := readproblem(*filename)
	var person user
	person.userinput()
	person.quiz(reader)
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

func (person *user) quiz(reader *csv.Reader) {
	i := 1
	var ans string
	for {
		problem, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Q", i, ": ", problem[0])
		fmt.Scanln(&ans)
		if ans == problem[1] {
			person.questionAnsweredCorrect = person.questionAnsweredCorrect + 1
		}
		if i == person.totalQuestionAttemted {
			break
		}
		i++
	}
}

func (person user) printResult() {
	fmt.Println(person.name, "attemted", person.totalQuestionAttemted, "Question and got",
		person.questionAnsweredCorrect, "answers correct")
}
