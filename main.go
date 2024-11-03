package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvfilename := flag.String("csv", "problem.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()
	// _ = csvfilename

	file, err := os.Open(*csvfilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open the CSV file:%s\n ", *csvfilename))
		// os.Exit(1)

	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file.")
	}
	// fmt.Println(lines)
	problems := parselines(lines)
	// fmt.Println(problems)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s=\n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
			fmt.Println("corrct")

		}
	}
	fmt.Printf("You scored %d out %d.\n", correct, len(problems))

}
func parselines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// go run main.go -csv="problem.csv"
