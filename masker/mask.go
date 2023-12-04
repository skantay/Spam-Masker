package mask

import (
	"fmt"
	"os"
)

type producer interface {
	produce()
}

type presenter interface {
	present()
}

type Service struct {
	prod producer
	pres presenter
}

type file struct {
	output, filepathFrom, filepathTo string
}

func (f *file) produce() error {
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

func (f *file) getFilePath() error {
	args := os.Args[1:]

	if len(args) > 2 || len(args) == 0 {
		return fmt.Errorf("Enter at least one file path. Maximum 2 file paths")
	}

	f.filepathFrom = args[0]

	if len(args) == 2 {
		f.filepathTo = args[1]
	}

	return nil
}

func (f *file) spamMasker() {
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

func (f *file) readFile() error {
	data, err := os.ReadFile(f.filepathFrom)
	if err != nil {
		return err
	}

	f.output = string(data)

	return nil
}

func (f *file) present() error {
	filepath := f.filepathFrom
	if f.filepathTo != "" {
		filepath = f.filepathTo
	}

	result, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Present error | present(): %w", err)
	}
	defer result.Close()

	buffer := []rune(f.output)

	for _, line := range buffer {
		_, err := result.WriteString(string(line))
		if err != nil {
			return fmt.Errorf("Present error | present(): %w", err)
		}
	}

	return nil
}

func Run() error {
	file := new(file)

	file.produce()

	err := file.present()
	if err != nil {
		return err
	}

	return nil
}
