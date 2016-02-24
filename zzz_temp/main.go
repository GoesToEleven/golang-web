package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			go func(m, n int) {
				fmt.Printf("%d %d\n", m, n)
			}(i, j)
		}
	}
	var a string
	fmt.Scanln(&a)
}
