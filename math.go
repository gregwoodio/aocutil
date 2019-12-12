package aocutil

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

// GreatestCommonDivisor finds the greatest common divisor between
// two ints.
func GreatestCommonDivisor(x, y int) int {
	gcd := 1

	absX := x
	if x < 0 {
		absX *= -1
	}
	absY := y
	if y < 0 {
		absY *= -1
	}

	for i := 1; i <= absX || i <= absY; i++ {
		if x%i == 0 && y%i == 0 {
			gcd = i
		}
	}

	return gcd
}
