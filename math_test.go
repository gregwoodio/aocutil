package aocutil

import "testing"

func TestPermutations(t *testing.T) {
	type testData struct {
		input    []int
		expected [][]int
	}

	testCases := []testData{
		testData{
			input: []int{1, 2},
			expected: [][]int{
				[]int{1, 2},
				[]int{2, 1},
			},
		},
		testData{
			input: []int{1, 2, 3},
			expected: [][]int{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
				[]int{2, 1, 3},
				[]int{2, 3, 1},
				[]int{3, 1, 2},
				[]int{3, 2, 1},
			},
		},
	}

	for _, td := range testCases {
		actual := Permutations(td.input)
		if len(actual) != len(td.expected) {
			t.Errorf("Expected %d but was %d", len(td.expected), len(actual))
		}
	}
}

func TestFactorial(t *testing.T) {
	type testData struct {
		input    uint64
		expected uint64
	}

	testCases := []testData{
		testData{
			input:    1,
			expected: 1,
		},
		testData{
			input:    2,
			expected: 2,
		},
		testData{
			input:    3,
			expected: 6,
		},
		testData{
			input:    4,
			expected: 24,
		},
		testData{
			input:    5,
			expected: 120,
		},
		testData{
			input:    6,
			expected: 720,
		},
		testData{
			input:    7,
			expected: 5040,
		},
	}

	for _, td := range testCases {
		actual := Factorial(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}

func TestGreatestCommonDivisor(t *testing.T) {
	type testData struct {
		inputs   []int
		expected int
	}

	testCases := []testData{
		testData{
			inputs:   []int{1, 2},
			expected: 1,
		},
		testData{
			inputs:   []int{8, 4},
			expected: 4,
		},
		testData{
			inputs:   []int{13, 52},
			expected: 13,
		},
		testData{
			inputs:   []int{3, 7},
			expected: 1,
		},
		testData{
			inputs:   []int{115, 65},
			expected: 5,
		},
		testData{
			inputs:   []int{3, 3},
			expected: 3,
		},
	}

	for _, td := range testCases {
		actual := GreatestCommonDivisor(td.inputs[0], td.inputs[1])

		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}
