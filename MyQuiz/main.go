package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "eazyQuestions.csv", "This is questions and answers.")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open the CSV file %s\n", *csvFilename)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse provided CSV file.")
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("[Question %d] %s = ", i, p.q)
		var answer string
		fmt.Scanf("%s", &answer)

		if answer == p.a {
			correct++
			fmt.Println("Correct!")
		} else {
			fmt.Println("----------------")
			fmt.Println("You are wrong...")
		}
	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problems))

}

type question struct {
	q string
	a string
}

func parseLines(lines [][]string) []question {
	ret := make([]question, len(lines))
	for i, line := range lines {
		ret[i] = question{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
