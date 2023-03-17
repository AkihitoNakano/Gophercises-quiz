package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "Read problems.csv data")

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Println("CSV file can't read")
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("CSV file doesn't read!")
		os.Exit(1)
	}

	parsedLines := lineParser(lines)

	var count int
	for i, p := range parsedLines {
		fmt.Printf("[question %d] %s = ", i, p.q)
		var answer string
		fmt.Scanf("%s", &answer)

		if answer == p.a {
			fmt.Println("You are correct!")
			count++
		}
	}
	fmt.Printf("Your score is %d out of %d", count, len(parsedLines))
}

type problems struct {
	q string
	a string
}

func lineParser(lines [][]string) []problems {
	// var p []problems
	p := make([]problems, len(lines))
	for i, line := range lines {
		p[i] = problems{
			q: line[0],
			a: line[1],
		}
	}
	return p
}
