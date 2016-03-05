package main

import "fmt"

func main() {
	var numone, numtwo int
	fmt.Print("Enter a number ")
	fmt.Scanf("%d \n", &numone)
	fmt.Print("Enter another number ")
	fmt.Scanf("%d \n", &numtwo)
	fmt.Println(numone)
	fmt.Println(numtwo)
}
