package mask

import "fmt"

type producer interface {
	produce() error
}

type presenter interface {
	present() error
}

type Service struct {
	prod producer
	pres presenter
}

type file struct {
	//nolint: structcheck
	output, filepathFrom, filepathTo string
}

func Run() error {
	text := new(file)
	prod := &fileProducer{text}
	pres := &filePresenter{text}

	service := &Service{
		prod: prod,
		pres: pres,
	}

	if err := service.prod.produce(); err != nil {
		return fmt.Errorf("err: run():%w", err)
	}

	if err := service.pres.present(); err != nil {
		return fmt.Errorf("err: run():%w", err)
	}

	return nil
}
