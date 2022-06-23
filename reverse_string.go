package reverse_string

func ReverseString(input string) (output string) {
	for _, r := range input {
		output = string(r) + output
	}

	return output
}
