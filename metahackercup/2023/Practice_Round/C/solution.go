package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(ca CaseArg) int {
	if ca.N == 1 {
		return 1
	}
	sort.Ints(ca.Weights)
	var sum int
	for _, num := range ca.Weights {
		sum += num
	}
	for i := len(ca.Weights) - 1; i >= 0; i-- {
		weight := ca.Weights[i]
		remain := sum - weight
		if remain%(ca.N-1) != 0 {
			continue
		}
		weight_per_day := remain / (ca.N - 1)
		if weight_per_day-weight <= 0 {
			continue
		}
		left, right := 0, len(ca.Weights)-1
		ast_found := true
		for left < right {
			if left == i {
				left++
			}
			if right == i {
				right--
			}
			if ca.Weights[left]+ca.Weights[right] != weight_per_day {
				ast_found = false
				break
			}
			left++
			right--
		}
		if ast_found {
			return weight_per_day - weight
		}
	}
	return -1
}

func runCases(cases []CaseArg) {
	for i := 0; i < len(cases); i++ {
		rst := solve(cases[i])
		ptext := fmt.Sprintf("Case #%d: %v", i+1, rst)
		fmt.Println(ptext)
		fmt.Fprintln(outputFile, ptext)
	}
}

func scan(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanCases() []CaseArg {
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)
	var count int = scanInt(scanner)
	var cases []CaseArg
	for i := 0; i < count; i++ {
		var ca CaseArg
		ca.N = scanInt(scanner)
		ca.Weights = make([]int, 2*ca.N-1)
		for i := 0; i < len(ca.Weights); i++ {
			ca.Weights[i] = scanInt(scanner)
		}
		cases = append(cases, ca)
	}
	return cases
}

type CaseArg struct {
	N       int
	Weights []int
}

var inputFile *os.File
var outputFile *os.File

var input string = "two_apples_a_day_input.txt"
var output string = "two_apples_a_day_output.txt"

func main() {
	inputFile, _ = os.Open(input)
	defer inputFile.Close()

	outputFile, _ = os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer outputFile.Close()

	cases := scanCases()
	runCases(cases)
}
