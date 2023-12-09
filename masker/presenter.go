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
