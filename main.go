package main

func spamMasker(buffer string) string {
	// Initialize variable
	var output []rune

	var toMask bool

	validate := "http://"

	input := []rune(buffer)

	for index := 0; index < len(input); index++ {
		// Check if last 7 chars of []output == http://
		if (len(output) >= len(validate)) && (string(output[index-len(validate):index]) == validate) {
			toMask = true
		} else if input[index] == ' ' {
			toMask = false
		}
		// Mask
		if toMask {
			output = append(output, '*')

			continue
		}

		char := input[index]
		output = append(output, char)
	}

	return string(output)
}
