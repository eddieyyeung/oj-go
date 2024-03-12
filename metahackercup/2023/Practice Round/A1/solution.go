package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputFile *os.File
var outputFile *os.File

func solve(ca CaseArg) bool {
	buns := (ca.S + ca.D) * 2
	patties := ca.S + ca.D*2
	needBuns := ca.K + 1
	needPatties := ca.K
	if buns < needBuns || patties < needPatties {
		return false
	}
	return true
}

func runCases(cases []CaseArg) {
	for i := 0; i < len(cases); i++ {
		var rst string
		if solve(cases[i]) {
			rst = "YES"
		} else {
			rst = "NO"
		}
		fmt.Printf("Case #%d: %s\n", i+1, rst)
		fmt.Fprintf(outputFile, "Case #%d: %s\n", i+1, rst)
	}
}

func scanCases() []CaseArg {
	scanner := bufio.NewScanner(inputFile)
	var count int
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &count)
	}
	var cases []CaseArg
	for i := 0; i < count; i++ {
		var ca CaseArg
		if scanner.Scan() {
			fmt.Sscanf(scanner.Text(), "%d %d %d", &ca.S, &ca.D, &ca.K)
			cases = append(cases, ca)
		}
	}
	return cases
}

type CaseArg struct {
	S int
	D int
	K int
}

var input string = "cheeseburger_corollary_1_input.txt"
var output string = "cheeseburger_corollary_1_output.txt"

func main() {
	inputFile, _ = os.Open(input)
	defer inputFile.Close()

	outputFile, _ = os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer outputFile.Close()

	cases := scanCases()
	runCases(cases)
}
