package aocutil

func reverse(str string) string {
	reversed := []byte{}

	for i := len(str) - 1; i >= 0; i-- {
		reversed = append(reversed, str[i])
	}

	return string(reversed)
}
