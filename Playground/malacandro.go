// This package is to determine the speed that a ship must
// go in order to reach the planet Malacandro in 28 days
package main

import "fmt"

func main () {
	
	const distance = 56000000
	const daysToMalacandro = 28

	fmt.Println((distance/daysToMalacandro), "kilometers/hour") // km/hr
}

