//How long does it take to get to Mars?
package main

import "fmt"

func main () {
	const lightSpeed = 299792 // km/s
	const hoursPerDay = 24
	var (
		distance = 56000000   // km
		speed = 100800)       // km/hr

	fmt.Println((distance/lightSpeed/60), "minutes")
	
	distance = 401000000
	fmt.Println((distance/lightSpeed/60), "minutes")

	distance = 96300000
	fmt.Println(distance/speed/hoursPerDay, "days")
}