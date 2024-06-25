package main

import (
	"fmt"
	"math/rand"
)
var era = "AD"

func main (){
	year := 2018

	switch month := rand.Intn(12) + 1; month {
	case 2:
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn
	}

}