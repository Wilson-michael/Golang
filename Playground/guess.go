// This is a program to guess a number between 1-100 that I've picked at random.
// It will also display if the guess is too large or too small.
// The number is 73
package main

import (
	"fmt"

	"math/rand"
)

func main() {
	var num = rand.Intn(100) + 1

	fmt.Println(num)

	if num == 73 {
		fmt.Println("Correct!")
	} else if num < 73 {
		fmt.Println("Too Small!")
	} else if num > 73 {
		fmt.Println("Too Large!")
	}
}
