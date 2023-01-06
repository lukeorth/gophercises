package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type problem struct {
    question    string
    answer    string
}

type quiz struct {
    answered    int
    score       int
    problems    []problem
}

func main() {
    // gracefully handle errors and exit
    var err error
    defer func() {
        if err != nil {
            log.Fatalln(err)
        }
    }()

    // get command args
    fname := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
    flag.Parse()

    // business logic
    q, err := loadQuiz(*fname)
    q.run()
}

func loadQuiz(name string) (quiz, error) {
    var q quiz

    // open file
    f, err := os.Open(name)
    if err != nil {
        return q, err
    }
    defer f.Close()

    // extract problems
    problems, err := parseProblems(f)
    if err != nil {
        return q, err
    }
    q.problems = problems

    return q, nil
}

func parseProblems(r io.Reader) ([]problem, error) {
    var problems []problem

    csvr := csv.NewReader(r)
    for {
        row, err := csvr.Read()
        if err != nil {
            if err == io.EOF {
                err = nil
            }
            return problems, err
        }
        problems = append(problems, problem{question: row[0], answer: row[1]})
    }
}

func (q *quiz) run() {
    s := bufio.NewScanner(os.Stdin)

    for i, p := range q.problems {
        fmt.Printf("Problem #%d: %s = ", i + 1, p.question)
        s.Scan()
        a := s.Text()
        if a == p.answer {
            q.score++
        }
        q.answered++
    }

    fmt.Printf("You scored %d out of %d.\n", q.score, len(q.problems))
}
