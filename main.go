package main

import (
	"fmt"
	mask "mymain/masker"
)

func main() {
	if err := mask.Run(); err != nil {
		fmt.Println(err)
	}
}
