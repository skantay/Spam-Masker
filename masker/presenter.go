package mask

import (
	"fmt"
	"os"
)

type filePresenter struct {
	*file
}

func (f *filePresenter) present() error {
	filepath := f.filepathFrom
	if f.filepathTo != "" {
		filepath = f.filepathTo
	}

	result, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("present error | present(): %w", err)
	}
	defer result.Close()

	var output string
	output = f.output
	for _, line := range output {
		_, err := result.WriteString(string(line))
		if err != nil {
			return fmt.Errorf("present error | present(): %w", err)
		}
	}

	return nil
}
