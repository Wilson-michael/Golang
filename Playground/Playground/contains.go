package main

import (
	"fmt"

	
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outside", 

	var exit = strings.Contains(command, "outside")

	fmt.Println("You leave the cave.", exit)

	var wearShades = true

	fmt.Println("The bright noonday sun is blinding to your eyes.", wearShades)

	var read = strings.Contains(command, "read")

	fmt.Println("You put on your shades, and realize you are at the top of a mountain overlooking a beautiful river valley. You look to your right and read a sign to the side of the cave at the head of a trail.", read)

	fmt.Println("The sign reads 'No minors'.")

	var age = 36
	var minor = age < 18

	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

}

