package mask

import (
	"errors"
	"fmt"
	"os"
)

type fileProducer struct {
	*file
}

func (f *fileProducer) produce() error {
	go f.spamMasker()

	if err := f.getFilePath(); err != nil {
		return fmt.Errorf("produce error | getFilePath(): %w", err)
	}

	if err := f.readFile(); err != nil {
		return fmt.Errorf("produce error | readFile(): %w", err)
	}

	return nil
}

func (f *fileProducer) getFilePath() error {
	//nolint: goerr113
	ErrGetFile := errors.New("enter at least one file path. Maximum 2 file paths")

	const TwoFiles int = 2

	args := os.Args[1:]

	if len(args) > 2 || len(args) == 0 {
		return ErrGetFile
	}

	f.filepathFrom = args[0]

	if len(args) == TwoFiles {
		f.filepathTo = args[1]
	}

	return nil
}

func (f *fileProducer) spamMasker() {
	// Initialize variable
	var output []rune

	var toMask bool

	validate := "http://"
	input := []rune(<-f.output)

	for index := 0; index < len(input); index++ {
		// Check if last 7 chars of []output == http://
		if (len(output) >= len(validate)) && (string(output[index-len(validate):index]) == validate) {
			toMask = true
		}

		// Check if link finished
		if input[index] == ' ' {
			toMask = false
		}

		// Mask
		if toMask {
			output = append(output, '*')
		} else {
			output = append(output, input[index])
		}
	}

	f.output <- string(output)
}

func (f *fileProducer) readFile() error {
	data, err := os.ReadFile(f.filepathFrom)
	if err != nil {
		return fmt.Errorf("os.ReadFile:%w", err)
	}

	f.output <- string(data)

	return nil
}
