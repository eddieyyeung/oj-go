package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func solve(ca CaseArg) int {
	var k1, k2, k3, k4 int = -1, -1, -1, -1

	k1 = 2*(ca.C/ca.B) - 1

	if ca.C >= ca.A {
		k2 = 1 + 2*((ca.C-ca.A)/ca.B)
	}
	if ca.C >= 2*ca.A {
		k3 = ca.C / ca.A
		k4 = 2 + 2*((ca.C-2*ca.A)/ca.B)
	}
	return max(k1, k2, k3, k4, 0)
}

func max(arr ...int) int {
	rst := math.MinInt
	for _, num := range arr {
		if num > rst {
			rst = num
		}
	}
	return rst
}

func runCases(cases []CaseArg) {
	for i := 0; i < len(cases); i++ {
		rst := solve(cases[i])
		ptext := fmt.Sprintf("Case #%d: %d", i+1, rst)
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
			fmt.Sscanf(scanner.Text(), "%d %d %d", &ca.A, &ca.B, &ca.C)
			cases = append(cases, ca)
		}
	}
	return cases
}

type CaseArg struct {
	A int
	B int
	C int
}

var inputFile *os.File
var outputFile *os.File

var input string = "cheeseburger_corollary_2_input.txt"
var output string = "cheeseburger_corollary_2_output.txt"

func main() {
	inputFile, _ = os.Open(input)
	defer inputFile.Close()

	outputFile, _ = os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer outputFile.Close()

	cases := scanCases()
	runCases(cases)
}
