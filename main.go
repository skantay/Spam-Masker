package main

func SpamMasker(buffer string) string {

	// Initialize variable
	var output []rune
	var toMask bool
	validate := "http://"
	input := []rune(buffer)

	for i := 0; i < len(input); i++ {


		// Check if last 7 chars of []output == http://
		if (len(output) >= len(validate)) && (string(output[i-len(validate):i]) == validate) {
			toMask = true
		}

		// Check if link finished
		if input[i] == ' ' {
			toMask = false
		}
		
		// Mask
		if toMask {
			output = append(output, '*')
		} else {
			output = append(output, input[i])
		}
	}

	return string(output)
}
