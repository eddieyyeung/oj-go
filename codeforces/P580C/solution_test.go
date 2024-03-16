package main

import (
	"bufio"
	"os"
	"testing"
)

func Test_solve(t *testing.T) {
	var inputFile *os.File
	var outputFile *os.File
	var input string = "case_input.txt"
	var output string = "case_output.txt"
	inputFile, _ = os.Open(input)
	defer inputFile.Close()
	outputFile, _ = os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer outputFile.Close()
	in := bufio.NewReader(inputFile)
	out := bufio.NewWriter(outputFile)
	defer out.Flush()

	for i := 0; i < 2; i++ {
		runCase(in, out)
	}
}
