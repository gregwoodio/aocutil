package aocutil

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadStringsFromFile reads the file and returns a []string
func ReadStringsFromFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Could not read file %s", filepath)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

// ReadIntsFromFile reads the file and returns a []int
func ReadIntsFromFile(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Could not read file %s", filepath)
	}
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

// ReadAndSplitInts reads a file line by line, and splits the line on the delimiter into ints.
func ReadAndSplitInts(filepath string, delimiter string) [][]int {
	file, _ := os.Open(filepath)
	reader := bufio.NewReader(file)

	var inputNumbers [][]int

	for line, _, _ := reader.ReadLine(); line != nil; {
		var numbers []int

		split := strings.Split(string(line), delimiter)

		for _, num := range split {
			float, _ := strconv.Atoi(num)
			numbers = append(numbers, float)
		}
		inputNumbers = append(inputNumbers, numbers)
		line, _, _ = reader.ReadLine()
	}

	return inputNumbers
}
