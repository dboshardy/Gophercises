package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := flag.Int("limit", 30, "the time limit for the quiz")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	timer := time.NewTimer(time.Second * time.Duration(*limit))
	numCorrect := int32(0)
	problems := readProblems(records)
Game:
	for i, prob := range problems {
		select {
		case <-timer.C:
			fmt.Println("Time's up!")
			break Game
		default:
			q := prob.question
			fmt.Printf("Question %d:\n", i+1)
			fmt.Printf("%s: ", q)
			var a string
			_, err := fmt.Scanf("%s\n", &a)
			if err != nil {
				_ = fmt.Errorf("Error processing answer!")
			}
			if a == prob.answer {
				numCorrect++
			}
		}
	}
	timer.Stop()
	fmt.Printf("You got %d of %d correct.", numCorrect, len(records))
}

type problem struct {
	question string
	answer   string
}

func readProblems(records [][]string) []problem {
	out := make([]problem, len(records))
	for i, record := range records {
		out[i] = problem{
			question: record[0],
			answer:   record[1],
		}
	}
	return out
}
