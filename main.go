package main


func SpamMasker(input string) string {
	var output []rune

	validate := "http://"
	var toMask bool
	buffer := []rune(input)

	for i := 0; i < len(buffer); i++ {
		if len(output) >= len(validate) {
			if string(output[i-len(validate):i]) == validate {
				toMask = true
			}
		}

		if buffer[i] == ' ' {
			toMask = false
		} else if toMask {
			char := '*'
			output = append(output, char)
			continue
		}

		output = append(output, buffer[i])
	}

	return string(output)
}
