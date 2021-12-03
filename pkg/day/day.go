package day

func copyInput(input []string) []string {
	copiedInput := make([]string, len(input))
	copy(copiedInput, input)
	return copiedInput
}
