package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(ca CaseArg) string {
	rst := ca.R > ca.C
	if rst {
		return "YES"
	}
	return "NO"
}

func runCases(cases []CaseArg) {
	for i := 0; i < len(cases); i++ {
		rst := solve(cases[i])
		ptext := fmt.Sprintf("Case #%d: %v", i+1, rst)
		fmt.Println(ptext)
		fmt.Fprintln(outputFile, ptext)
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
			fmt.Sscanf(scanner.Text(), "%d %d %d %d", &ca.R, &ca.C, &ca.A, &ca.B)
			cases = append(cases, ca)
		}
	}
	return cases
}

type CaseArg struct {
	R int
	C int
	A int
	B int
}

var inputFile *os.File
var outputFile *os.File

var input string = "dim_sum_delivery_input.txt"
var output string = "dim_sum_delivery_output.txt"

func main() {
	inputFile, _ = os.Open(input)
	defer inputFile.Close()

	outputFile, _ = os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer outputFile.Close()

	cases := scanCases()
	runCases(cases)
}
