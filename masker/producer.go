package mask

import (
	"fmt"
	"os"
)

type fileProducer struct {
	*file
}

func (f *fileProducer) produce() error {
	err := f.getFilePath()
	if err != nil {
		return fmt.Errorf("Produce error | getFilePath(): %w", err)
	}

	err = f.readFile()
	if err != nil {
		return fmt.Errorf("Produce error | readFile(): %w", err)
	}

	f.spamMasker()
	return nil
}

func (f *fileProducer) getFilePath() error {
	args := os.Args[1:]

	if len(args) > 2 || len(args) == 0 {
		return fmt.Errorf("enter at least one file path. Maximum 2 file paths")
	}

	f.filepathFrom = args[0]

	if len(args) == 2 {
		f.filepathTo = args[1]
	}

	return nil
}

func (f *fileProducer) spamMasker() {
	// Initialize variable
	var output []rune
	var toMask bool
	validate := "http://"
	input := []rune(f.output)

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

	f.output = string(output)
}

func (f *fileProducer) readFile() error {
	data, err := os.ReadFile(f.filepathFrom)
	if err != nil {
		return err
	}

	f.output = string(data)

	return nil
}
