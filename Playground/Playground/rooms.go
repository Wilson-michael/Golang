package main

import "fmt"

func main () {
	var room = "cave" 
	
	if room == "cave" {
		fmt.Println("You find yourself in a dimly lit cavern, filled with cobwebs and dust.")
	} else if room == "entrance" {
		fmt.Println("There is a cavern entrance and a path to the east, heading up the mountain.")
	} else if room == "mountain" {
		fmt.Println("There is a cliff here. A path leads west down the mountain.")
	} else {
		fmt.Println("Everything else is white.")
	}
}