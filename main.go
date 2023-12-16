package main

import (
	"fmt"
	"time"

	//nolint: depguard
	mask "github.com/skantay/Spam-Masker/masker"
)

func main() {
	startTime := time.Now()
	if err := mask.Run(); err != nil {
		//nolint: forbidigo
		fmt.Println(err)
	}
	endTime := time.Now()

	execT := endTime.Sub(startTime)
	fmt.Printf("Function executed in %v\n", execT)
}
