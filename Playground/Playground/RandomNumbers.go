package main

import (
	"fmt"

	"math/rand"
)

func main () { // this number is a random distance between
			   // earth and mars
	var num = rand.Intn(401000000) + 56000000
	
	fmt.Println(num)

}