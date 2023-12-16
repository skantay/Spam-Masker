package mask

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/adhocore/chin"
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

func (f *fileProducer) spamMasker(wg *sync.WaitGroup, inputB string) {
	// Initialize variable
	var output []rune

	var toMask bool

	validate := "http://"
<<<<<<< HEAD
	input := []rune(<-f.output)
=======
	input := []rune(inputB)
>>>>>>> new

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

<<<<<<< HEAD
	f.output <- string(output)
=======
	f.output = f.output + "\n" + string(output)
	wg.Done()
>>>>>>> new
}

func (f *fileProducer) readFile() error {
	file, err := os.Open(f.filepathFrom)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	wg := new(sync.WaitGroup)

	s := chin.New().WithWait(wg)
	go s.Start()

	for scanner.Scan() {
		wg.Add(1)
		go f.spamMasker(wg, scanner.Text())
	}
	s.Stop()

<<<<<<< HEAD
	f.output <- string(data)
=======
	wg.Wait()
>>>>>>> new

	return nil
}
