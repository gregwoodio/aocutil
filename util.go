package aocutil

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

// Permutations takes the given slice of int and returns all permutations
// of it.
func Permutations(input []int) [][]int {
	inputCopy := make([]int, len(input))
	copy(inputCopy, input)
	output := [][]int{
		inputCopy,
	}

	limit := int(Factorial(uint64(len(input))))
	n := len(input) - 1
	var i, j int

	for c := 1; c < limit; c++ {
		i = n - 1
		j = n

		for input[i] > input[i+1] {
			i--
		}

		for input[j] < input[i] {
			j--
		}

		input[i], input[j] = input[j], input[i]

		j = n
		i++

		for i < j {
			input[i], input[j] = input[j], input[i]
			i++
			j--
		}

		inputCopy = make([]int, len(input))
		copy(inputCopy, input)
		output = append(output, inputCopy)
	}

	return output
}

// Factorial computes factorial of the given number.
// ie. 3! = 3 * 2 * 1
func Factorial(input uint64) uint64 {
	if input > 0 {
		return input * Factorial(input-1)
	}

	return 1
}
