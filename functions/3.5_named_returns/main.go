package main

import "fmt"

func yearsUntilEvent(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	return 18 - age, 21 - age, 25 - age
}

func yearsUntilEventImplicitReturn(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	yearsUntilDrinking = 21 - age
	yearsUntilCarRental = 25 - age
	return
}

func main() {
	yUA, yUD, yUCR := yearsUntilEventImplicitReturn(15)
	fmt.Println("Years until adult:", yUA, "Years until drinking:", yUD, "Years until car rental:", yUCR)
}
