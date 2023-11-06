package mask

import (
	"fmt"
	"log"
	"os"
)

type producer interface {
	produce()
}

type presenter interface {
	present()
}

type service struct {
	prod producer
	pres presenter
}

type worker struct {
	service
	text, from, to string
	result         []string
}

func Run() {
	worker := new(worker)
	result, err := worker.produce()
	if err != nil {
		log.Fatal(err)
	}

	worker.result = result

	_, err = worker.present()
	if err != nil {
		log.Fatal(err)
	}
}

func (w *worker) getText() {

	err := w.getFilePath()
	if err != nil {
		log.Fatal(err)
	}

	err = w.getString()
	if err != nil {
		log.Fatal(err)
	}
}

func (w *worker) produce() ([]string, error) {
	w.getText()

	text := w.text

	maskedText := w.spamMasker(text)

	result := make([]string, 0, 1)
	result = append(result, maskedText)

	return result, nil
}

func (w *worker) present() ([]string, error) {
	filepath := w.from
	if w.to != "" {
		filepath = w.to
	}

	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, line := range w.result {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil, nil
}

func (w *worker) getString() error {

	data, err := os.ReadFile(w.from)
	if err != nil {
		return err
	}

	w.text = string(data)

	return nil
}

func (w *worker) getFilePath() error {
	args := os.Args[1:]

	if len(args) > 2 || len(args) == 0 {
		return fmt.Errorf("Enter at least one file path. Maximum 2 file paths")
	}

	w.from = args[0]

	if len(args) == 2 {
		w.to = args[1]
	}

	return nil
}

func (s *service) spamMasker(buffer string) string {
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
