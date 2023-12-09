package main

import (
	"fmt"

	//nolint: depguard
	mask "github.com/skantay/Spam-Masker/masker"
)

func main() {
	if err := mask.Run(); err != nil {
		//nolint: forbidigo
		fmt.Println(err)
	}
}
