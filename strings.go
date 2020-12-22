package aocutil

// Reverse the provided string
func Reverse(str string) string {
	reversed := []byte{}

	for i := len(str) - 1; i >= 0; i-- {
		reversed = append(reversed, str[i])
	}

	return string(reversed)
}
